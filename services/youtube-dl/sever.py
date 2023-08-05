from concurrent.futures import ThreadPoolExecutor
import json
import logging
from typing import TypeVar
import grpc

import yt_dlp

from youtube_dl_pb2 import GetVideoInfoResponse, Response
from error_pb2 import ErrorResponse
from youtube_dl_pb2_grpc import (
    GetVideoInfoServiceServicer,
    add_GetVideoInfoServiceServicer_to_server,
)

import google.protobuf.json_format as json_format


def get_video_info(videoURL: str):
    info = None
    ydl_opts = {}
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        try:
            info = ydl.extract_info(videoURL, download=False)
        except:
            logging.critical("error in extracting info")
    return info


# https://stackoverflow.com/questions/60345426/json-to-protobuf-in-python
# its important to enable ignore_unkwon_fields because
# 'data'  object contains more fields that what was specified from the message
# having this set to true will allow to remove those unknown fields
T = TypeVar("T")


def to_json(data, type):
    data_type = type
    json_format.Parse(
        json.dumps(data, indent=None), data_type, ignore_unknown_fields=True
    )
    return data_type


def from_json(data):
    return json_format.MessageToJson(data)


class GetVideoInfoService(GetVideoInfoServiceServicer):
    def GetVideo(self, request, context):
        data = get_video_info(videoURL=request.videoURL)
        resp = Response()
        if data is not None:
            result = to_json(data=data, type=GetVideoInfoResponse())
            # When an error message shows up with the follow:
            # Assignment not allowed to message, map, or repeated field "field_name" in protocol message object
            # use CopyFrom
            # https://stackoverflow.com/questions/18376190/attributeerror-assignment-not-allowed-to-composite-field-task-in-protocol-mes
            resp.info_response.CopyFrom(result)
        else:
            errObj = {"error_message": "video not found", "status_code": 404}
            errorMsg = to_json(data=errObj, type=ErrorResponse())
            resp.error.CopyFrom(errorMsg)
            # context.set_code(grpc.StatusCode.NOT_FOUND)

        return resp


if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor())
    add_GetVideoInfoServiceServicer_to_server(GetVideoInfoService(), server)
    port = 9999
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    logging.info("server ready on port %r", port)
    server.wait_for_termination()

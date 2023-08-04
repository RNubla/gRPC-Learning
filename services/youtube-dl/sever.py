from concurrent.futures import ThreadPoolExecutor
import json
import logging
import grpc

import yt_dlp

from youtube_dl_pb2 import GetVideoInfoResponse

from youtube_dl_pb2_grpc import (
    GetVideoInfoServiceServicer,
    add_GetVideoInfoServiceServicer_to_server,
)

import google.protobuf.json_format as json_format


# https://stackoverflow.com/questions/60345426/json-to-protobuf-in-python
# its important to enable ignore_unkwon_fields because
# 'data'  object contains more fields that what was specified from the message
# having this set to true will allow to remove those unknown fields
def toJson(data, type):
    return json_format.Parse(
        json.dumps(data, indent=None), type(), ignore_unknown_fields=True
    )


def get_video_info(videoURL: str):
    info = None
    ydl_opts = {}
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        try:
            info = ydl.extract_info(videoURL, download=False)
        except:
            logging.critical("error in extracting info")
    return info


class GetVideoInfoService(GetVideoInfoServiceServicer):
    def GetVideo(self, request, context):
        data = get_video_info(videoURL=request.videoURL)
        resp = toJson(data=data, type=GetVideoInfoResponse)
        # print(resp)
        return resp


if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor())
    add_GetVideoInfoServiceServicer_to_server(GetVideoInfoService(), server)
    port = 9999
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    logging.info("server ready on port %r", port)
    server.wait_for_termination()

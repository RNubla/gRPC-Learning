from concurrent.futures import ThreadPoolExecutor
import logging
import grpc

import yt_dlp
from automapper import mapper
from typing import List


from youtube_dl_pb2 import GetVideoInfoResponse
from youtube_dl_pb2_grpc import (
    GetVideoInfoServiceServicer,
    add_GetVideoInfoServiceServicer_to_server,
)


class RequestedFormatModel:
    def __init__(
        self,
        filesize: int,
        format_id: str,
        url: str,
        ext: str,
        video_ext: str,
        audio_ext: str,
    ):
        self.filesize = filesize
        self.format_id = format_id
        self.url = url
        self.ext = ext
        self.video_ext = video_ext
        self.audio_ext = audio_ext


class GetVideoInfoResponseModel:
    def __init__(
        self,
        id: str,
        title: str,
        thumbnail: str,
        requested_formats: List[RequestedFormatModel],
    ):
        self.id = id
        self.title = title
        self.thumbnail = thumbnail
        self.requested_formats = requested_formats


def get_video_info(videoURL: str):
    info = None
    ydl_opts = {}
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        try:
            info = ydl.extract_info(videoURL, download=False)
        except:
            logging.critical("error in extracting info")
            # raise HTTPException(status_code=404)
            # return HTTPException(status_code=404)
    info = mapper.to(GetVideoInfoResponseModel).map(info)
    info.requested_formats = mapper.to(RequestedFormatModel).map(info.requested_formats)
    return info


class GetVideoInfoService(GetVideoInfoServiceServicer):
    def GetVideo(self, request, context):
        data = get_video_info(videoURL=request.videoURL)
        print(data.requested_formats)
        # logging.debug(data)
        resp = GetVideoInfoResponse(
            id=data.id,
            title=data.title,
            thumbnail=data.thumbnail,
            requested_formats=data.requested_formats,
        )
        # resp = GetVideoInfoResponse()
        # logging.debug(inspect.signature(GetVideoInfoResponse))
        return resp


if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor())
    add_GetVideoInfoServiceServicer_to_server(GetVideoInfoService(), server)
    port = 9999
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    logging.info("server ready on port %r", port)
    server.wait_for_termination()

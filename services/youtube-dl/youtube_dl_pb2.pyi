from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetVideoInfoRequest(_message.Message):
    __slots__ = ["videoURL"]
    VIDEOURL_FIELD_NUMBER: _ClassVar[int]
    videoURL: str
    def __init__(self, videoURL: _Optional[str] = ...) -> None: ...

class RequestedFormats(_message.Message):
    __slots__ = ["filesize", "format_id", "url", "ext", "video_ext", "audio_ext"]
    FILESIZE_FIELD_NUMBER: _ClassVar[int]
    FORMAT_ID_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    EXT_FIELD_NUMBER: _ClassVar[int]
    VIDEO_EXT_FIELD_NUMBER: _ClassVar[int]
    AUDIO_EXT_FIELD_NUMBER: _ClassVar[int]
    filesize: int
    format_id: str
    url: str
    ext: str
    video_ext: str
    audio_ext: str
    def __init__(self, filesize: _Optional[int] = ..., format_id: _Optional[str] = ..., url: _Optional[str] = ..., ext: _Optional[str] = ..., video_ext: _Optional[str] = ..., audio_ext: _Optional[str] = ...) -> None: ...

class GetVideoInfoResponse(_message.Message):
    __slots__ = ["id", "title", "thumbnail", "requested_formats"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    THUMBNAIL_FIELD_NUMBER: _ClassVar[int]
    REQUESTED_FORMATS_FIELD_NUMBER: _ClassVar[int]
    id: str
    title: str
    thumbnail: str
    requested_formats: _containers.RepeatedCompositeFieldContainer[RequestedFormats]
    def __init__(self, id: _Optional[str] = ..., title: _Optional[str] = ..., thumbnail: _Optional[str] = ..., requested_formats: _Optional[_Iterable[_Union[RequestedFormats, _Mapping]]] = ...) -> None: ...

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/gpu/services.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.common import metrics_pb2 as datahub_dot_common_dot_metrics__pb2
from datahub.common import queries_pb2 as datahub_dot_common_dot_queries__pb2
from datahub.gpu import gpu_pb2 as datahub_dot_gpu_dot_gpu__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/gpu/services.proto',
  package='prophetstor.datahub.gpu',
  syntax='proto3',
  serialized_options=b'Z\037prophetstor.com/api/datahub/gpu',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1a\x64\x61tahub/gpu/services.proto\x12\x17prophetstor.datahub.gpu\x1a\x1c\x64\x61tahub/common/metrics.proto\x1a\x1c\x64\x61tahub/common/queries.proto\x1a\x15\x64\x61tahub/gpu/gpu.proto\x1a\x17google/rpc/status.proto\"z\n\x0fListGpusRequest\x12\x43\n\x0fquery_condition\x18\x01 \x01(\x0b\x32*.prophetstor.datahub.common.QueryCondition\x12\x0c\n\x04host\x18\x02 \x01(\t\x12\x14\n\x0cminor_number\x18\x03 \x01(\t\"b\n\x10ListGpusResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12*\n\x04gpus\x18\x02 \x03(\x0b\x32\x1c.prophetstor.datahub.gpu.Gpu\"\xbe\x01\n\x15ListGpuMetricsRequest\x12\x43\n\x0fquery_condition\x18\x01 \x01(\x0b\x32*.prophetstor.datahub.common.QueryCondition\x12<\n\x0cmetric_types\x18\x02 \x03(\x0e\x32&.prophetstor.datahub.common.MetricType\x12\x0c\n\x04host\x18\x03 \x01(\t\x12\x14\n\x0cminor_number\x18\x04 \x01(\t\"u\n\x16ListGpuMetricsResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x37\n\x0bgpu_metrics\x18\x02 \x03(\x0b\x32\".prophetstor.datahub.gpu.GpuMetric\"^\n\x1b\x43reateGpuPredictionsRequest\x12?\n\x0fgpu_predictions\x18\x01 \x03(\x0b\x32&.prophetstor.datahub.gpu.GpuPrediction\"\xc2\x01\n\x19ListGpuPredictionsRequest\x12\x43\n\x0fquery_condition\x18\x01 \x01(\x0b\x32*.prophetstor.datahub.common.QueryCondition\x12\x0c\n\x04host\x18\x02 \x01(\t\x12\x14\n\x0cminor_number\x18\x03 \x01(\t\x12\x13\n\x0bgranularity\x18\x04 \x01(\x03\x12\x10\n\x08model_id\x18\x05 \x01(\t\x12\x15\n\rprediction_id\x18\x06 \x01(\t\"\x81\x01\n\x1aListGpuPredictionsResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12?\n\x0fgpu_predictions\x18\x02 \x03(\x0b\x32&.prophetstor.datahub.gpu.GpuPredictionB!Z\x1fprophetstor.com/api/datahub/gpub\x06proto3'
  ,
  dependencies=[datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,datahub_dot_common_dot_queries__pb2.DESCRIPTOR,datahub_dot_gpu_dot_gpu__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_LISTGPUSREQUEST = _descriptor.Descriptor(
  name='ListGpusRequest',
  full_name='prophetstor.datahub.gpu.ListGpusRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='prophetstor.datahub.gpu.ListGpusRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='host', full_name='prophetstor.datahub.gpu.ListGpusRequest.host', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='minor_number', full_name='prophetstor.datahub.gpu.ListGpusRequest.minor_number', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=163,
  serialized_end=285,
)


_LISTGPUSRESPONSE = _descriptor.Descriptor(
  name='ListGpusResponse',
  full_name='prophetstor.datahub.gpu.ListGpusResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.datahub.gpu.ListGpusResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gpus', full_name='prophetstor.datahub.gpu.ListGpusResponse.gpus', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=287,
  serialized_end=385,
)


_LISTGPUMETRICSREQUEST = _descriptor.Descriptor(
  name='ListGpuMetricsRequest',
  full_name='prophetstor.datahub.gpu.ListGpuMetricsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='prophetstor.datahub.gpu.ListGpuMetricsRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_types', full_name='prophetstor.datahub.gpu.ListGpuMetricsRequest.metric_types', index=1,
      number=2, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='host', full_name='prophetstor.datahub.gpu.ListGpuMetricsRequest.host', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='minor_number', full_name='prophetstor.datahub.gpu.ListGpuMetricsRequest.minor_number', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=388,
  serialized_end=578,
)


_LISTGPUMETRICSRESPONSE = _descriptor.Descriptor(
  name='ListGpuMetricsResponse',
  full_name='prophetstor.datahub.gpu.ListGpuMetricsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.datahub.gpu.ListGpuMetricsResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gpu_metrics', full_name='prophetstor.datahub.gpu.ListGpuMetricsResponse.gpu_metrics', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=580,
  serialized_end=697,
)


_CREATEGPUPREDICTIONSREQUEST = _descriptor.Descriptor(
  name='CreateGpuPredictionsRequest',
  full_name='prophetstor.datahub.gpu.CreateGpuPredictionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='gpu_predictions', full_name='prophetstor.datahub.gpu.CreateGpuPredictionsRequest.gpu_predictions', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=699,
  serialized_end=793,
)


_LISTGPUPREDICTIONSREQUEST = _descriptor.Descriptor(
  name='ListGpuPredictionsRequest',
  full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='host', full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest.host', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='minor_number', full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest.minor_number', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='granularity', full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest.granularity', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='model_id', full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest.model_id', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='prediction_id', full_name='prophetstor.datahub.gpu.ListGpuPredictionsRequest.prediction_id', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=796,
  serialized_end=990,
)


_LISTGPUPREDICTIONSRESPONSE = _descriptor.Descriptor(
  name='ListGpuPredictionsResponse',
  full_name='prophetstor.datahub.gpu.ListGpuPredictionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.datahub.gpu.ListGpuPredictionsResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gpu_predictions', full_name='prophetstor.datahub.gpu.ListGpuPredictionsResponse.gpu_predictions', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=993,
  serialized_end=1122,
)

_LISTGPUSREQUEST.fields_by_name['query_condition'].message_type = datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_LISTGPUSRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTGPUSRESPONSE.fields_by_name['gpus'].message_type = datahub_dot_gpu_dot_gpu__pb2._GPU
_LISTGPUMETRICSREQUEST.fields_by_name['query_condition'].message_type = datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_LISTGPUMETRICSREQUEST.fields_by_name['metric_types'].enum_type = datahub_dot_common_dot_metrics__pb2._METRICTYPE
_LISTGPUMETRICSRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTGPUMETRICSRESPONSE.fields_by_name['gpu_metrics'].message_type = datahub_dot_gpu_dot_gpu__pb2._GPUMETRIC
_CREATEGPUPREDICTIONSREQUEST.fields_by_name['gpu_predictions'].message_type = datahub_dot_gpu_dot_gpu__pb2._GPUPREDICTION
_LISTGPUPREDICTIONSREQUEST.fields_by_name['query_condition'].message_type = datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_LISTGPUPREDICTIONSRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTGPUPREDICTIONSRESPONSE.fields_by_name['gpu_predictions'].message_type = datahub_dot_gpu_dot_gpu__pb2._GPUPREDICTION
DESCRIPTOR.message_types_by_name['ListGpusRequest'] = _LISTGPUSREQUEST
DESCRIPTOR.message_types_by_name['ListGpusResponse'] = _LISTGPUSRESPONSE
DESCRIPTOR.message_types_by_name['ListGpuMetricsRequest'] = _LISTGPUMETRICSREQUEST
DESCRIPTOR.message_types_by_name['ListGpuMetricsResponse'] = _LISTGPUMETRICSRESPONSE
DESCRIPTOR.message_types_by_name['CreateGpuPredictionsRequest'] = _CREATEGPUPREDICTIONSREQUEST
DESCRIPTOR.message_types_by_name['ListGpuPredictionsRequest'] = _LISTGPUPREDICTIONSREQUEST
DESCRIPTOR.message_types_by_name['ListGpuPredictionsResponse'] = _LISTGPUPREDICTIONSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListGpusRequest = _reflection.GeneratedProtocolMessageType('ListGpusRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTGPUSREQUEST,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.ListGpusRequest)
  })
_sym_db.RegisterMessage(ListGpusRequest)

ListGpusResponse = _reflection.GeneratedProtocolMessageType('ListGpusResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTGPUSRESPONSE,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.ListGpusResponse)
  })
_sym_db.RegisterMessage(ListGpusResponse)

ListGpuMetricsRequest = _reflection.GeneratedProtocolMessageType('ListGpuMetricsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTGPUMETRICSREQUEST,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.ListGpuMetricsRequest)
  })
_sym_db.RegisterMessage(ListGpuMetricsRequest)

ListGpuMetricsResponse = _reflection.GeneratedProtocolMessageType('ListGpuMetricsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTGPUMETRICSRESPONSE,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.ListGpuMetricsResponse)
  })
_sym_db.RegisterMessage(ListGpuMetricsResponse)

CreateGpuPredictionsRequest = _reflection.GeneratedProtocolMessageType('CreateGpuPredictionsRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEGPUPREDICTIONSREQUEST,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.CreateGpuPredictionsRequest)
  })
_sym_db.RegisterMessage(CreateGpuPredictionsRequest)

ListGpuPredictionsRequest = _reflection.GeneratedProtocolMessageType('ListGpuPredictionsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTGPUPREDICTIONSREQUEST,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.ListGpuPredictionsRequest)
  })
_sym_db.RegisterMessage(ListGpuPredictionsRequest)

ListGpuPredictionsResponse = _reflection.GeneratedProtocolMessageType('ListGpuPredictionsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTGPUPREDICTIONSRESPONSE,
  '__module__' : 'datahub.gpu.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.gpu.ListGpuPredictionsResponse)
  })
_sym_db.RegisterMessage(ListGpuPredictionsResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
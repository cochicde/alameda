# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/data/types.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.common import metrics_pb2 as datahub_dot_common_dot_metrics__pb2
from datahub.common import rawdata_pb2 as datahub_dot_common_dot_rawdata__pb2
from datahub.common import types_pb2 as datahub_dot_common_dot_types__pb2
from datahub.schemas import types_pb2 as datahub_dot_schemas_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/data/types.proto',
  package='prophetstor.datahub.data',
  syntax='proto3',
  serialized_options=b'Z prophetstor.com/api/datahub/data',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x18\x64\x61tahub/data/types.proto\x12\x18prophetstor.datahub.data\x1a\x1c\x64\x61tahub/common/metrics.proto\x1a\x1c\x64\x61tahub/common/rawdata.proto\x1a\x1a\x64\x61tahub/common/types.proto\x1a\x1b\x64\x61tahub/schemas/types.proto\"x\n\x04\x44\x61ta\x12<\n\x0bschema_meta\x18\x01 \x01(\x0b\x32\'.prophetstor.datahub.schemas.SchemaMeta\x12\x32\n\x07rawdata\x18\x02 \x03(\x0b\x32!.prophetstor.datahub.data.Rawdata\"\x9a\x02\n\x07Rawdata\x12\x13\n\x0bmeasurement\x18\x01 \x01(\t\x12;\n\x0bmetric_type\x18\x02 \x01(\x0e\x32&.prophetstor.datahub.common.MetricType\x12G\n\x11resource_boundary\x18\x03 \x01(\x0e\x32,.prophetstor.datahub.common.ResourceBoundary\x12\x41\n\x0eresource_quota\x18\x04 \x01(\x0e\x32).prophetstor.datahub.common.ResourceQuota\x12\x31\n\x06groups\x18\x05 \x03(\x0b\x32!.prophetstor.datahub.common.GroupB\"Z prophetstor.com/api/datahub/datab\x06proto3'
  ,
  dependencies=[datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,datahub_dot_common_dot_types__pb2.DESCRIPTOR,datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,])




_DATA = _descriptor.Descriptor(
  name='Data',
  full_name='prophetstor.datahub.data.Data',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.datahub.data.Data.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='rawdata', full_name='prophetstor.datahub.data.Data.rawdata', index=1,
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
  serialized_start=171,
  serialized_end=291,
)


_RAWDATA = _descriptor.Descriptor(
  name='Rawdata',
  full_name='prophetstor.datahub.data.Rawdata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='measurement', full_name='prophetstor.datahub.data.Rawdata.measurement', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='prophetstor.datahub.data.Rawdata.metric_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_boundary', full_name='prophetstor.datahub.data.Rawdata.resource_boundary', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_quota', full_name='prophetstor.datahub.data.Rawdata.resource_quota', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='groups', full_name='prophetstor.datahub.data.Rawdata.groups', index=4,
      number=5, type=11, cpp_type=10, label=3,
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
  serialized_start=294,
  serialized_end=576,
)

_DATA.fields_by_name['schema_meta'].message_type = datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_DATA.fields_by_name['rawdata'].message_type = _RAWDATA
_RAWDATA.fields_by_name['metric_type'].enum_type = datahub_dot_common_dot_metrics__pb2._METRICTYPE
_RAWDATA.fields_by_name['resource_boundary'].enum_type = datahub_dot_common_dot_types__pb2._RESOURCEBOUNDARY
_RAWDATA.fields_by_name['resource_quota'].enum_type = datahub_dot_common_dot_types__pb2._RESOURCEQUOTA
_RAWDATA.fields_by_name['groups'].message_type = datahub_dot_common_dot_rawdata__pb2._GROUP
DESCRIPTOR.message_types_by_name['Data'] = _DATA
DESCRIPTOR.message_types_by_name['Rawdata'] = _RAWDATA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Data = _reflection.GeneratedProtocolMessageType('Data', (_message.Message,), {
  'DESCRIPTOR' : _DATA,
  '__module__' : 'datahub.data.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.data.Data)
  })
_sym_db.RegisterMessage(Data)

Rawdata = _reflection.GeneratedProtocolMessageType('Rawdata', (_message.Message,), {
  'DESCRIPTOR' : _RAWDATA,
  '__module__' : 'datahub.data.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.data.Rawdata)
  })
_sym_db.RegisterMessage(Rawdata)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

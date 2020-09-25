# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/metrics/metrics.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.common import metrics_pb2 as datahub_dot_common_dot_metrics__pb2
from datahub.common import rawdata_pb2 as datahub_dot_common_dot_rawdata__pb2
from datahub.resources import metadata_pb2 as datahub_dot_resources_dot_metadata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/metrics/metrics.proto',
  package='prophetstor.datahub.metrics',
  syntax='proto3',
  serialized_options=b'Z#prophetstor.com/api/datahub/metrics',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1d\x64\x61tahub/metrics/metrics.proto\x12\x1bprophetstor.datahub.metrics\x1a\x1c\x64\x61tahub/common/metrics.proto\x1a\x1c\x64\x61tahub/common/rawdata.proto\x1a datahub/resources/metadata.proto\"\\\n\x0f\x43ontainerMetric\x12\x0c\n\x04name\x18\x01 \x01(\t\x12;\n\x0bmetric_data\x18\x02 \x03(\x0b\x32&.prophetstor.datahub.common.MetricData\"\x94\x01\n\tPodMetric\x12>\n\x0bobject_meta\x18\x01 \x01(\x0b\x32).prophetstor.datahub.resources.ObjectMeta\x12G\n\x11\x63ontainer_metrics\x18\x02 \x03(\x0b\x32,.prophetstor.datahub.metrics.ContainerMetric\"\xc2\x01\n\x10\x43ontrollerMetric\x12>\n\x0bobject_meta\x18\x01 \x01(\x0b\x32).prophetstor.datahub.resources.ObjectMeta\x12\x31\n\x04kind\x18\x02 \x01(\x0e\x32#.prophetstor.datahub.resources.Kind\x12;\n\x0bmetric_data\x18\x03 \x03(\x0b\x32&.prophetstor.datahub.common.MetricData\"\x90\x01\n\x11\x41pplicationMetric\x12>\n\x0bobject_meta\x18\x01 \x01(\x0b\x32).prophetstor.datahub.resources.ObjectMeta\x12;\n\x0bmetric_data\x18\x02 \x03(\x0b\x32&.prophetstor.datahub.common.MetricData\"\x8e\x01\n\x0fNamespaceMetric\x12>\n\x0bobject_meta\x18\x01 \x01(\x0b\x32).prophetstor.datahub.resources.ObjectMeta\x12;\n\x0bmetric_data\x18\x02 \x03(\x0b\x32&.prophetstor.datahub.common.MetricData\"\x89\x01\n\nNodeMetric\x12>\n\x0bobject_meta\x18\x01 \x01(\x0b\x32).prophetstor.datahub.resources.ObjectMeta\x12;\n\x0bmetric_data\x18\x02 \x03(\x0b\x32&.prophetstor.datahub.common.MetricData\"\x8c\x01\n\rClusterMetric\x12>\n\x0bobject_meta\x18\x01 \x01(\x0b\x32).prophetstor.datahub.resources.ObjectMeta\x12;\n\x0bmetric_data\x18\x02 \x03(\x0b\x32&.prophetstor.datahub.common.MetricData\"\x85\x01\n\x0bWriteMetric\x12;\n\x0bmetric_type\x18\x01 \x01(\x0e\x32&.prophetstor.datahub.common.MetricType\x12\x39\n\nwrite_data\x18\x02 \x01(\x0b\x32%.prophetstor.datahub.common.WriteDataB%Z#prophetstor.com/api/datahub/metricsb\x06proto3'
  ,
  dependencies=[datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,datahub_dot_resources_dot_metadata__pb2.DESCRIPTOR,])




_CONTAINERMETRIC = _descriptor.Descriptor(
  name='ContainerMetric',
  full_name='prophetstor.datahub.metrics.ContainerMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='prophetstor.datahub.metrics.ContainerMetric.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='prophetstor.datahub.metrics.ContainerMetric.metric_data', index=1,
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
  serialized_start=156,
  serialized_end=248,
)


_PODMETRIC = _descriptor.Descriptor(
  name='PodMetric',
  full_name='prophetstor.datahub.metrics.PodMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.datahub.metrics.PodMetric.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='container_metrics', full_name='prophetstor.datahub.metrics.PodMetric.container_metrics', index=1,
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
  serialized_start=251,
  serialized_end=399,
)


_CONTROLLERMETRIC = _descriptor.Descriptor(
  name='ControllerMetric',
  full_name='prophetstor.datahub.metrics.ControllerMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.datahub.metrics.ControllerMetric.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='kind', full_name='prophetstor.datahub.metrics.ControllerMetric.kind', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='prophetstor.datahub.metrics.ControllerMetric.metric_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
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
  serialized_start=402,
  serialized_end=596,
)


_APPLICATIONMETRIC = _descriptor.Descriptor(
  name='ApplicationMetric',
  full_name='prophetstor.datahub.metrics.ApplicationMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.datahub.metrics.ApplicationMetric.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='prophetstor.datahub.metrics.ApplicationMetric.metric_data', index=1,
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
  serialized_start=599,
  serialized_end=743,
)


_NAMESPACEMETRIC = _descriptor.Descriptor(
  name='NamespaceMetric',
  full_name='prophetstor.datahub.metrics.NamespaceMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.datahub.metrics.NamespaceMetric.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='prophetstor.datahub.metrics.NamespaceMetric.metric_data', index=1,
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
  serialized_start=746,
  serialized_end=888,
)


_NODEMETRIC = _descriptor.Descriptor(
  name='NodeMetric',
  full_name='prophetstor.datahub.metrics.NodeMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.datahub.metrics.NodeMetric.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='prophetstor.datahub.metrics.NodeMetric.metric_data', index=1,
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
  serialized_start=891,
  serialized_end=1028,
)


_CLUSTERMETRIC = _descriptor.Descriptor(
  name='ClusterMetric',
  full_name='prophetstor.datahub.metrics.ClusterMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.datahub.metrics.ClusterMetric.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='prophetstor.datahub.metrics.ClusterMetric.metric_data', index=1,
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
  serialized_start=1031,
  serialized_end=1171,
)


_WRITEMETRIC = _descriptor.Descriptor(
  name='WriteMetric',
  full_name='prophetstor.datahub.metrics.WriteMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='prophetstor.datahub.metrics.WriteMetric.metric_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_data', full_name='prophetstor.datahub.metrics.WriteMetric.write_data', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1174,
  serialized_end=1307,
)

_CONTAINERMETRIC.fields_by_name['metric_data'].message_type = datahub_dot_common_dot_metrics__pb2._METRICDATA
_PODMETRIC.fields_by_name['object_meta'].message_type = datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_PODMETRIC.fields_by_name['container_metrics'].message_type = _CONTAINERMETRIC
_CONTROLLERMETRIC.fields_by_name['object_meta'].message_type = datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_CONTROLLERMETRIC.fields_by_name['kind'].enum_type = datahub_dot_resources_dot_metadata__pb2._KIND
_CONTROLLERMETRIC.fields_by_name['metric_data'].message_type = datahub_dot_common_dot_metrics__pb2._METRICDATA
_APPLICATIONMETRIC.fields_by_name['object_meta'].message_type = datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_APPLICATIONMETRIC.fields_by_name['metric_data'].message_type = datahub_dot_common_dot_metrics__pb2._METRICDATA
_NAMESPACEMETRIC.fields_by_name['object_meta'].message_type = datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_NAMESPACEMETRIC.fields_by_name['metric_data'].message_type = datahub_dot_common_dot_metrics__pb2._METRICDATA
_NODEMETRIC.fields_by_name['object_meta'].message_type = datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_NODEMETRIC.fields_by_name['metric_data'].message_type = datahub_dot_common_dot_metrics__pb2._METRICDATA
_CLUSTERMETRIC.fields_by_name['object_meta'].message_type = datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_CLUSTERMETRIC.fields_by_name['metric_data'].message_type = datahub_dot_common_dot_metrics__pb2._METRICDATA
_WRITEMETRIC.fields_by_name['metric_type'].enum_type = datahub_dot_common_dot_metrics__pb2._METRICTYPE
_WRITEMETRIC.fields_by_name['write_data'].message_type = datahub_dot_common_dot_rawdata__pb2._WRITEDATA
DESCRIPTOR.message_types_by_name['ContainerMetric'] = _CONTAINERMETRIC
DESCRIPTOR.message_types_by_name['PodMetric'] = _PODMETRIC
DESCRIPTOR.message_types_by_name['ControllerMetric'] = _CONTROLLERMETRIC
DESCRIPTOR.message_types_by_name['ApplicationMetric'] = _APPLICATIONMETRIC
DESCRIPTOR.message_types_by_name['NamespaceMetric'] = _NAMESPACEMETRIC
DESCRIPTOR.message_types_by_name['NodeMetric'] = _NODEMETRIC
DESCRIPTOR.message_types_by_name['ClusterMetric'] = _CLUSTERMETRIC
DESCRIPTOR.message_types_by_name['WriteMetric'] = _WRITEMETRIC
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerMetric = _reflection.GeneratedProtocolMessageType('ContainerMetric', (_message.Message,), {
  'DESCRIPTOR' : _CONTAINERMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.ContainerMetric)
  })
_sym_db.RegisterMessage(ContainerMetric)

PodMetric = _reflection.GeneratedProtocolMessageType('PodMetric', (_message.Message,), {
  'DESCRIPTOR' : _PODMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.PodMetric)
  })
_sym_db.RegisterMessage(PodMetric)

ControllerMetric = _reflection.GeneratedProtocolMessageType('ControllerMetric', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLERMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.ControllerMetric)
  })
_sym_db.RegisterMessage(ControllerMetric)

ApplicationMetric = _reflection.GeneratedProtocolMessageType('ApplicationMetric', (_message.Message,), {
  'DESCRIPTOR' : _APPLICATIONMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.ApplicationMetric)
  })
_sym_db.RegisterMessage(ApplicationMetric)

NamespaceMetric = _reflection.GeneratedProtocolMessageType('NamespaceMetric', (_message.Message,), {
  'DESCRIPTOR' : _NAMESPACEMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.NamespaceMetric)
  })
_sym_db.RegisterMessage(NamespaceMetric)

NodeMetric = _reflection.GeneratedProtocolMessageType('NodeMetric', (_message.Message,), {
  'DESCRIPTOR' : _NODEMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.NodeMetric)
  })
_sym_db.RegisterMessage(NodeMetric)

ClusterMetric = _reflection.GeneratedProtocolMessageType('ClusterMetric', (_message.Message,), {
  'DESCRIPTOR' : _CLUSTERMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.ClusterMetric)
  })
_sym_db.RegisterMessage(ClusterMetric)

WriteMetric = _reflection.GeneratedProtocolMessageType('WriteMetric', (_message.Message,), {
  'DESCRIPTOR' : _WRITEMETRIC,
  '__module__' : 'datahub.metrics.metrics_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.metrics.WriteMetric)
  })
_sym_db.RegisterMessage(WriteMetric)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

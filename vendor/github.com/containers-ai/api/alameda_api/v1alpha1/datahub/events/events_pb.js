// source: alameda_api/v1alpha1/datahub/events/events.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var alameda_api_v1alpha1_datahub_events_types_pb = require('../../../../alameda_api/v1alpha1/datahub/events/types_pb.js');
goog.object.extend(proto, alameda_api_v1alpha1_datahub_events_types_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.events.Event', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.events.Event, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.events.Event.displayName = 'proto.containersai.alameda.v1alpha1.datahub.events.Event';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.events.Event.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.Event} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.toObject = function(includeInstance, msg) {
  var f, obj = {
    time: (f = msg.getTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    id: jspb.Message.getFieldWithDefault(msg, 2, ""),
    clusterId: jspb.Message.getFieldWithDefault(msg, 3, ""),
    source: (f = msg.getSource()) && alameda_api_v1alpha1_datahub_events_types_pb.EventSource.toObject(includeInstance, f),
    type: jspb.Message.getFieldWithDefault(msg, 5, 0),
    version: jspb.Message.getFieldWithDefault(msg, 6, 0),
    level: jspb.Message.getFieldWithDefault(msg, 7, 0),
    subject: (f = msg.getSubject()) && alameda_api_v1alpha1_datahub_events_types_pb.K8SObjectReference.toObject(includeInstance, f),
    message: jspb.Message.getFieldWithDefault(msg, 9, ""),
    data: jspb.Message.getFieldWithDefault(msg, 10, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.events.Event;
  return proto.containersai.alameda.v1alpha1.datahub.events.Event.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.Event} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTime(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setClusterId(value);
      break;
    case 4:
      var value = new alameda_api_v1alpha1_datahub_events_types_pb.EventSource;
      reader.readMessage(value,alameda_api_v1alpha1_datahub_events_types_pb.EventSource.deserializeBinaryFromReader);
      msg.setSource(value);
      break;
    case 5:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.events.EventType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 6:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.events.EventVersion} */ (reader.readEnum());
      msg.setVersion(value);
      break;
    case 7:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.events.EventLevel} */ (reader.readEnum());
      msg.setLevel(value);
      break;
    case 8:
      var value = new alameda_api_v1alpha1_datahub_events_types_pb.K8SObjectReference;
      reader.readMessage(value,alameda_api_v1alpha1_datahub_events_types_pb.K8SObjectReference.deserializeBinaryFromReader);
      msg.setSubject(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.events.Event.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.Event} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTime();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getClusterId();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getSource();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      alameda_api_v1alpha1_datahub_events_types_pb.EventSource.serializeBinaryToWriter
    );
  }
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
  f = message.getVersion();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getLevel();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
  f = message.getSubject();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      alameda_api_v1alpha1_datahub_events_types_pb.K8SObjectReference.serializeBinaryToWriter
    );
  }
  f = message.getMessage();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getData();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
};


/**
 * optional google.protobuf.Timestamp time = 1;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 1));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setTime = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.clearTime = function() {
  return this.setTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.hasTime = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string id = 2;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string cluster_id = 3;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getClusterId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setClusterId = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional EventSource source = 4;
 * @return {?proto.containersai.alameda.v1alpha1.datahub.events.EventSource}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getSource = function() {
  return /** @type{?proto.containersai.alameda.v1alpha1.datahub.events.EventSource} */ (
    jspb.Message.getWrapperField(this, alameda_api_v1alpha1_datahub_events_types_pb.EventSource, 4));
};


/**
 * @param {?proto.containersai.alameda.v1alpha1.datahub.events.EventSource|undefined} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setSource = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.clearSource = function() {
  return this.setSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.hasSource = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional EventType type = 5;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.EventType}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getType = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.events.EventType} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.EventType} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
};


/**
 * optional EventVersion version = 6;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.EventVersion}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getVersion = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.events.EventVersion} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.EventVersion} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setVersion = function(value) {
  return jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * optional EventLevel level = 7;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.EventLevel}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getLevel = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.events.EventLevel} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.EventLevel} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setLevel = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};


/**
 * optional K8SObjectReference subject = 8;
 * @return {?proto.containersai.alameda.v1alpha1.datahub.events.K8SObjectReference}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getSubject = function() {
  return /** @type{?proto.containersai.alameda.v1alpha1.datahub.events.K8SObjectReference} */ (
    jspb.Message.getWrapperField(this, alameda_api_v1alpha1_datahub_events_types_pb.K8SObjectReference, 8));
};


/**
 * @param {?proto.containersai.alameda.v1alpha1.datahub.events.K8SObjectReference|undefined} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setSubject = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.clearSubject = function() {
  return this.setSubject(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.hasSubject = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string message = 9;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setMessage = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string data = 10;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.getData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.events.Event} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.events.Event.prototype.setData = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


goog.object.extend(exports, proto.containersai.alameda.v1alpha1.datahub.events);
// source: datahub/schemas/schemas.proto
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

var datahub_schemas_types_pb = require('../../datahub/schemas/types_pb.js');
goog.object.extend(proto, datahub_schemas_types_pb);
goog.exportSymbol('proto.prophetstor.datahub.schemas.Schema', null, global);
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
proto.prophetstor.datahub.schemas.Schema = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.datahub.schemas.Schema.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.datahub.schemas.Schema, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.schemas.Schema.displayName = 'proto.prophetstor.datahub.schemas.Schema';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.datahub.schemas.Schema.repeatedFields_ = [2];



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
proto.prophetstor.datahub.schemas.Schema.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.schemas.Schema.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.schemas.Schema} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.schemas.Schema.toObject = function(includeInstance, msg) {
  var f, obj = {
    schemaMeta: (f = msg.getSchemaMeta()) && datahub_schemas_types_pb.SchemaMeta.toObject(includeInstance, f),
    measurementsList: jspb.Message.toObjectList(msg.getMeasurementsList(),
    datahub_schemas_types_pb.Measurement.toObject, includeInstance)
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
 * @return {!proto.prophetstor.datahub.schemas.Schema}
 */
proto.prophetstor.datahub.schemas.Schema.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.schemas.Schema;
  return proto.prophetstor.datahub.schemas.Schema.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.schemas.Schema} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.schemas.Schema}
 */
proto.prophetstor.datahub.schemas.Schema.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new datahub_schemas_types_pb.SchemaMeta;
      reader.readMessage(value,datahub_schemas_types_pb.SchemaMeta.deserializeBinaryFromReader);
      msg.setSchemaMeta(value);
      break;
    case 2:
      var value = new datahub_schemas_types_pb.Measurement;
      reader.readMessage(value,datahub_schemas_types_pb.Measurement.deserializeBinaryFromReader);
      msg.addMeasurements(value);
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
proto.prophetstor.datahub.schemas.Schema.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.schemas.Schema.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.schemas.Schema} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.schemas.Schema.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSchemaMeta();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      datahub_schemas_types_pb.SchemaMeta.serializeBinaryToWriter
    );
  }
  f = message.getMeasurementsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      datahub_schemas_types_pb.Measurement.serializeBinaryToWriter
    );
  }
};


/**
 * optional SchemaMeta schema_meta = 1;
 * @return {?proto.prophetstor.datahub.schemas.SchemaMeta}
 */
proto.prophetstor.datahub.schemas.Schema.prototype.getSchemaMeta = function() {
  return /** @type{?proto.prophetstor.datahub.schemas.SchemaMeta} */ (
    jspb.Message.getWrapperField(this, datahub_schemas_types_pb.SchemaMeta, 1));
};


/**
 * @param {?proto.prophetstor.datahub.schemas.SchemaMeta|undefined} value
 * @return {!proto.prophetstor.datahub.schemas.Schema} returns this
*/
proto.prophetstor.datahub.schemas.Schema.prototype.setSchemaMeta = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.datahub.schemas.Schema} returns this
 */
proto.prophetstor.datahub.schemas.Schema.prototype.clearSchemaMeta = function() {
  return this.setSchemaMeta(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.datahub.schemas.Schema.prototype.hasSchemaMeta = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Measurement measurements = 2;
 * @return {!Array<!proto.prophetstor.datahub.schemas.Measurement>}
 */
proto.prophetstor.datahub.schemas.Schema.prototype.getMeasurementsList = function() {
  return /** @type{!Array<!proto.prophetstor.datahub.schemas.Measurement>} */ (
    jspb.Message.getRepeatedWrapperField(this, datahub_schemas_types_pb.Measurement, 2));
};


/**
 * @param {!Array<!proto.prophetstor.datahub.schemas.Measurement>} value
 * @return {!proto.prophetstor.datahub.schemas.Schema} returns this
*/
proto.prophetstor.datahub.schemas.Schema.prototype.setMeasurementsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.datahub.schemas.Measurement=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.datahub.schemas.Measurement}
 */
proto.prophetstor.datahub.schemas.Schema.prototype.addMeasurements = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.datahub.schemas.Measurement, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.datahub.schemas.Schema} returns this
 */
proto.prophetstor.datahub.schemas.Schema.prototype.clearMeasurementsList = function() {
  return this.setMeasurementsList([]);
};


goog.object.extend(exports, proto.prophetstor.datahub.schemas);

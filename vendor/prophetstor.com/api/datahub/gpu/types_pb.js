// source: datahub/gpu/types.proto
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

goog.exportSymbol('proto.prophetstor.datahub.gpu.GpuMetadata', null, global);
goog.exportSymbol('proto.prophetstor.datahub.gpu.GpuSpec', null, global);
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
proto.prophetstor.datahub.gpu.GpuMetadata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.gpu.GpuMetadata, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.gpu.GpuMetadata.displayName = 'proto.prophetstor.datahub.gpu.GpuMetadata';
}
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
proto.prophetstor.datahub.gpu.GpuSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.gpu.GpuSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.gpu.GpuSpec.displayName = 'proto.prophetstor.datahub.gpu.GpuSpec';
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
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.gpu.GpuMetadata.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.gpu.GpuMetadata} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.gpu.GpuMetadata.toObject = function(includeInstance, msg) {
  var f, obj = {
    host: jspb.Message.getFieldWithDefault(msg, 1, ""),
    instance: jspb.Message.getFieldWithDefault(msg, 2, ""),
    job: jspb.Message.getFieldWithDefault(msg, 3, ""),
    minorNumber: jspb.Message.getFieldWithDefault(msg, 4, "")
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
 * @return {!proto.prophetstor.datahub.gpu.GpuMetadata}
 */
proto.prophetstor.datahub.gpu.GpuMetadata.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.gpu.GpuMetadata;
  return proto.prophetstor.datahub.gpu.GpuMetadata.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.gpu.GpuMetadata} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.gpu.GpuMetadata}
 */
proto.prophetstor.datahub.gpu.GpuMetadata.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setHost(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setInstance(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setJob(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setMinorNumber(value);
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
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.gpu.GpuMetadata.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.gpu.GpuMetadata} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.gpu.GpuMetadata.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHost();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getInstance();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getJob();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getMinorNumber();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string host = 1;
 * @return {string}
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.getHost = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.gpu.GpuMetadata} returns this
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.setHost = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string instance = 2;
 * @return {string}
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.getInstance = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.gpu.GpuMetadata} returns this
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.setInstance = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string job = 3;
 * @return {string}
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.getJob = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.gpu.GpuMetadata} returns this
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.setJob = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string minor_number = 4;
 * @return {string}
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.getMinorNumber = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.gpu.GpuMetadata} returns this
 */
proto.prophetstor.datahub.gpu.GpuMetadata.prototype.setMinorNumber = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};





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
proto.prophetstor.datahub.gpu.GpuSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.gpu.GpuSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.gpu.GpuSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.gpu.GpuSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    memoryTotal: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0)
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
 * @return {!proto.prophetstor.datahub.gpu.GpuSpec}
 */
proto.prophetstor.datahub.gpu.GpuSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.gpu.GpuSpec;
  return proto.prophetstor.datahub.gpu.GpuSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.gpu.GpuSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.gpu.GpuSpec}
 */
proto.prophetstor.datahub.gpu.GpuSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setMemoryTotal(value);
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
proto.prophetstor.datahub.gpu.GpuSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.gpu.GpuSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.gpu.GpuSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.gpu.GpuSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMemoryTotal();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
};


/**
 * optional float memory_total = 1;
 * @return {number}
 */
proto.prophetstor.datahub.gpu.GpuSpec.prototype.getMemoryTotal = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.prophetstor.datahub.gpu.GpuSpec} returns this
 */
proto.prophetstor.datahub.gpu.GpuSpec.prototype.setMemoryTotal = function(value) {
  return jspb.Message.setProto3FloatField(this, 1, value);
};


goog.object.extend(exports, proto.prophetstor.datahub.gpu);

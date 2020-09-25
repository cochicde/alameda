// source: datahub/weavescope/services.proto
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

var google_rpc_status_pb = require('../../google/rpc/status_pb.js');
goog.object.extend(proto, google_rpc_status_pb);
goog.exportSymbol('proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest', null, global);
goog.exportSymbol('proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest', null, global);
goog.exportSymbol('proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest', null, global);
goog.exportSymbol('proto.prophetstor.datahub.weavescope.WeaveScopeResponse', null, global);
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
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.displayName = 'proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest';
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
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.displayName = 'proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest';
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
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.displayName = 'proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest';
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
proto.prophetstor.datahub.weavescope.WeaveScopeResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.weavescope.WeaveScopeResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.weavescope.WeaveScopeResponse.displayName = 'proto.prophetstor.datahub.weavescope.WeaveScopeResponse';
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
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    hostId: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest;
  return proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setHostId(value);
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
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHostId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string host_id = 1;
 * @return {string}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.prototype.getHostId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest} returns this
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeHostsRequest.prototype.setHostId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    podId: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest;
  return proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setPodId(value);
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
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPodId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string pod_id = 1;
 * @return {string}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.prototype.getPodId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest} returns this
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopePodsRequest.prototype.setPodId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    containerId: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest;
  return proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setContainerId(value);
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
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContainerId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string container_id = 1;
 * @return {string}
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.prototype.getContainerId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest} returns this
 */
proto.prophetstor.datahub.weavescope.ListWeaveScopeContainersRequest.prototype.setContainerId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.weavescope.WeaveScopeResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    status: (f = msg.getStatus()) && google_rpc_status_pb.Status.toObject(includeInstance, f),
    rawdata: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse}
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.weavescope.WeaveScopeResponse;
  return proto.prophetstor.datahub.weavescope.WeaveScopeResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse}
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_rpc_status_pb.Status;
      reader.readMessage(value,google_rpc_status_pb.Status.deserializeBinaryFromReader);
      msg.setStatus(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRawdata(value);
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
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.weavescope.WeaveScopeResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStatus();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_rpc_status_pb.Status.serializeBinaryToWriter
    );
  }
  f = message.getRawdata();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional google.rpc.Status status = 1;
 * @return {?proto.google.rpc.Status}
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.getStatus = function() {
  return /** @type{?proto.google.rpc.Status} */ (
    jspb.Message.getWrapperField(this, google_rpc_status_pb.Status, 1));
};


/**
 * @param {?proto.google.rpc.Status|undefined} value
 * @return {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse} returns this
*/
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.setStatus = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse} returns this
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.clearStatus = function() {
  return this.setStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.hasStatus = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string rawdata = 2;
 * @return {string}
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.getRawdata = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.weavescope.WeaveScopeResponse} returns this
 */
proto.prophetstor.datahub.weavescope.WeaveScopeResponse.prototype.setRawdata = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.prophetstor.datahub.weavescope);

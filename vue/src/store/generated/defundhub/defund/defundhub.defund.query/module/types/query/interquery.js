/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "defundhub.defund.query";
const baseInterquery = {
    index: "",
    height: "",
    path: "",
    chainId: "",
    typeName: "",
    creator: "",
};
export const Interquery = {
    encode(message, writer = Writer.create()) {
        if (message.index !== "") {
            writer.uint32(10).string(message.index);
        }
        if (message.height !== "") {
            writer.uint32(18).string(message.height);
        }
        if (message.path !== "") {
            writer.uint32(26).string(message.path);
        }
        if (message.chainId !== "") {
            writer.uint32(34).string(message.chainId);
        }
        if (message.typeName !== "") {
            writer.uint32(42).string(message.typeName);
        }
        if (message.creator !== "") {
            writer.uint32(50).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseInterquery };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.index = reader.string();
                    break;
                case 2:
                    message.height = reader.string();
                    break;
                case 3:
                    message.path = reader.string();
                    break;
                case 4:
                    message.chainId = reader.string();
                    break;
                case 5:
                    message.typeName = reader.string();
                    break;
                case 6:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseInterquery };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        if (object.height !== undefined && object.height !== null) {
            message.height = String(object.height);
        }
        else {
            message.height = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = String(object.path);
        }
        else {
            message.path = "";
        }
        if (object.chainId !== undefined && object.chainId !== null) {
            message.chainId = String(object.chainId);
        }
        else {
            message.chainId = "";
        }
        if (object.typeName !== undefined && object.typeName !== null) {
            message.typeName = String(object.typeName);
        }
        else {
            message.typeName = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        message.height !== undefined && (obj.height = message.height);
        message.path !== undefined && (obj.path = message.path);
        message.chainId !== undefined && (obj.chainId = message.chainId);
        message.typeName !== undefined && (obj.typeName = message.typeName);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseInterquery };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        if (object.height !== undefined && object.height !== null) {
            message.height = object.height;
        }
        else {
            message.height = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = object.path;
        }
        else {
            message.path = "";
        }
        if (object.chainId !== undefined && object.chainId !== null) {
            message.chainId = object.chainId;
        }
        else {
            message.chainId = "";
        }
        if (object.typeName !== undefined && object.typeName !== null) {
            message.typeName = object.typeName;
        }
        else {
            message.typeName = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};

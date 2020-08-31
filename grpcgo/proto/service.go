syntax = "proto3";

package proto;

message Request {
  int64 a = 1;
  int64 b = 2;
}

message Response { 
  int64 result = 1; 
}

service AddService {
  rpc Division(Request) returns (Response);
  rpc Multiply(Request) returns (Response);
}
// Copyright (c) 2017 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

namespace java com.uber.dosa

typedef map<string, Value> FieldValueMap

enum ElemType {
   BOOL,
   BLOB,
   STRING,
   INT32,
   INT64,
   DOUBLE,
   TIMESTAMP,
   UUID,

   // Reserve a few enum types in case we want to support few more
   RESERVED0,
   RESERVED1,
   RESERVED2,
   RESERVED3,
}

union RawValue {
 1: optional binary binaryValue                 // BLOB, UUID
 2: optional bool boolValue                     // BOOL
 3: optional double doubleValue                 // DOUBLE
 4: optional i32 int32Value                     // INT32
 5: optional i64 (js.type = "Long") int64Value  // INT64, TIMESTAMP
 6: optional string stringValue                 // STRING
}

// make it union in case we want to support collections like set
union Value {
   1: optional RawValue elemValue
}

struct SchemaRef {
    1: optional string scope
    2: optional string namePrefix
    3: optional string entityName
    4: optional i32 version
}

struct FieldTag {
   1: optional string name
   2: optional string value
}

struct FieldDesc {
   1: optional ElemType type
   2: optional set<FieldTag> tags
}

struct ClusteringKey {
   1: optional string name
   2: optional bool asc
}

struct PrimaryKey {
   1: optional list<string> partitionKeys
   2: optional list<ClusteringKey> clusteringKeys
}

struct IndexDefinition {
   1: optional PrimaryKey key
}

struct EntityDefinition {
   1: optional string name
   2: optional map<string, FieldDesc> fieldDescs
   3: optional PrimaryKey primaryKey
   4: optional map<string, IndexDefinition> Indexes
   5: optional bool enableETL
}

struct Error {
   1: optional i32 errCode
   2: optional string msg
   3: optional bool shouldRetry
}

struct CreateRequest {
   1: optional SchemaRef ref
   2: optional FieldValueMap entityValues
}

struct ReadRequest {
   1: optional SchemaRef ref
   2: optional FieldValueMap keyValues
   3: optional set<string> fieldsToRead
}

struct ReadResponse {
   1: optional FieldValueMap entityValues
}

struct MultiReadRequest {
   1: optional SchemaRef ref
   2: optional list<FieldValueMap> keyValues
   3: optional set<string> fieldsToRead
}

union EntityOrError{
   1: optional FieldValueMap entityValues
   2: optional Error error
}

struct MultiReadResponse {
   1: optional list<EntityOrError> results
}

struct MultiUpsertResponse {
   1: optional list<Error> errors
}

struct UpsertRequest {
    1: optional SchemaRef ref
    2: optional FieldValueMap entityValues
}

struct MultiUpsertRequest {
    1: optional SchemaRef ref
    2: optional list<FieldValueMap> entities
}

struct RemoveRequest {
   1: optional SchemaRef ref
   2: optional FieldValueMap keyValues
}

struct MultiRemoveRequest {
   1: optional SchemaRef ref 
   2: optional list<FieldValueMap> keyValues
}

struct MultiRemoveResponse {
   1: optional list<Error> errors
}

enum Operator {
   EQ,
   LT,
   GT,
   LT_OR_EQ,
   GT_OR_EQ,
}

struct Field {
   1: optional string name
   2: optional Value value
}

struct Condition {
   1: optional Operator op
   2: optional Field field
}

struct RangeRequest {
   1: optional SchemaRef ref
   2: optional string token
   3: optional i32 limit
   4: optional list<Condition> conditions
   5: optional set<string> fieldsToRead
}

struct RangeResponse {
   1: optional list<FieldValueMap> entities
   2: optional string nextToken
}

struct RemoveRangeRequest {
   1: optional SchemaRef ref
   2: optional list<Condition> conditions
}

struct SearchRequest {
   1: optional SchemaRef ref
   2: optional string token
   3: optional i32 limit
   4: optional Field searchBy
   5: optional set<string> fieldsToRead
}

struct SearchResponse {
   1: optional list<FieldValueMap> entities
   2: optional string nextToken
}

struct ScanRequest {
   1: optional SchemaRef ref
   2: optional string token
   3: optional i32 limit
   4: optional set<string> fieldsToRead
}

struct ScanResponse {
   1: optional list<FieldValueMap> entities
   2: optional string nextToken
}

struct CanUpsertSchemaRequest {
   1: optional string scope
   2: optional string namePrefix
   3: optional list<EntityDefinition> entityDefs
}

struct CanUpsertSchemaResponse {
   // latest upserted version
   1: optional i32 version
}

struct CheckSchemaRequest {
   1: optional string scope
   2: optional string namePrefix
   3: optional list<EntityDefinition> entityDefs
}

struct CheckSchemaResponse {
   1: optional i32 version
}

struct UpsertSchemaRequest {
   1: optional string scope
   2: optional string namePrefix
   3: optional list<EntityDefinition> entityDefs
   4: optional bool dryRun
}

struct UpsertSchemaResponse {
   1: optional i32 version
   2: optional string status
}

struct CheckSchemaStatusRequest {
   1: optional string scope
   2: optional string namePrefix
   3: optional i32 version
}

struct CheckSchemaStatusResponse {
    1: optional i32 version
    2: optional string status
}

struct CreateScopeRequest {
   1: optional string name
}

struct ScopeExistsRequest {
   1: optional string name
}

struct ScopeExistsResponse {
   1: optional bool exists
}

struct TruncateScopeRequest {
   1: optional string name
}
struct DropScopeRequest {
   1: optional string name
}

exception BadRequestError {
    1: required string err
    2: optional string message
    3: optional i32 errorCode
}

exception InternalServerError {
    1: required string err
    2: optional string message
    3: optional i32 errorCode
}

exception BadSchemaError {
    // EntityName -> Error msg
    1: required map<string, string> reasons
}

service Dosa {
   void createIfNotExists(
       1: CreateRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   ReadResponse read (
       1: ReadRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   MultiReadResponse multiRead (
       1: MultiReadRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   void upsert (
       1: UpsertRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   MultiUpsertResponse multiUpsert (
       1: MultiUpsertRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )
   
   void remove (
       1: RemoveRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   MultiRemoveResponse multiRemove (
       1: MultiRemoveRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   RangeResponse range (
       1: RangeRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   void removeRange (
       1: RemoveRangeRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   SearchResponse search (
       1: SearchRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   ScanResponse scan (
       1: ScanRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   CanUpsertSchemaResponse canUpsertSchema(
       1: CanUpsertSchemaRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
       3: BadSchemaError schemaError
   )

   CheckSchemaResponse checkSchema(
       1: CheckSchemaRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
       3: BadSchemaError schemaError
   )

   UpsertSchemaResponse upsertSchema(
       1: UpsertSchemaRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
       3: BadSchemaError schemaError
   )

   CheckSchemaStatusResponse checkSchemaStatus(
       1: CheckSchemaStatusRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   void createScope(
       1: CreateScopeRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   ScopeExistsResponse scopeExists(
       1: ScopeExistsRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   void truncateScope(
       1: TruncateScopeRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )

   void dropScope(
       1: DropScopeRequest request
   ) throws (
       1: BadRequestError clientError
       2: InternalServerError serverError
   )
}


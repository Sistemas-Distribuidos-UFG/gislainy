///
//  Generated code. Do not modify.
//  source: questions/questions.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'questions.pb.dart' as $0;
export 'questions.pb.dart';

class QuestionsClient extends $grpc.Client {
  static final _$calculateReajustSalary =
      $grpc.ClientMethod<$0.Employee, $0.Employee>(
          '/questions.Questions/CalculateReajustSalary',
          ($0.Employee value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Employee.fromBuffer(value));

  QuestionsClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.Employee> calculateReajustSalary($0.Employee request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$calculateReajustSalary, request,
        options: options);
  }
}

abstract class QuestionsServiceBase extends $grpc.Service {
  $core.String get $name => 'questions.Questions';

  QuestionsServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Employee, $0.Employee>(
        'CalculateReajustSalary',
        calculateReajustSalary_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Employee.fromBuffer(value),
        ($0.Employee value) => value.writeToBuffer()));
  }

  $async.Future<$0.Employee> calculateReajustSalary_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Employee> request) async {
    return calculateReajustSalary(call, await request);
  }

  $async.Future<$0.Employee> calculateReajustSalary(
      $grpc.ServiceCall call, $0.Employee request);
}

import 'dart:ffi';

import 'package:questions_app/common/grpc_commons.dart';
import 'package:questions_app/model/questions/questions.pb.dart';
import 'package:questions_app/model/questions/questions.pbgrpc.dart';

class QuestionsService {
  static Future<Employee> calculateReajustSalary(
      {String name, String role, double salary}) async {
    var stub = QuestionsClient(GrpcClientSingleton().client);
    final employee = Employee()
      ..name = name
      ..role = role
      ..salary = salary;
    return await stub.calculateReajustSalary(employee);
  }
}

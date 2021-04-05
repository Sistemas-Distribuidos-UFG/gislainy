import 'package:flutter/material.dart';
import 'package:questions_app/model/questions/questions.pb.dart';
import 'package:questions_app/services/questions_service.dart';

class QuestionView extends StatefulWidget {
  @override
  _QuestionViewState createState() => _QuestionViewState();
}

class _QuestionViewState extends State<QuestionView> {
  Employee employee = Employee();

  TextEditingController nameController = TextEditingController();
  TextEditingController salaryController = TextEditingController();
  TextEditingController roleController = TextEditingController();

  double salaryReajusted = 0.0;

  @override
  void initState() {
    employee.name = "Gislainy";
    employee.role = "Developer";
    employee.salary = 2000;
    // TODO: implement initState
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Questions Distribud Systems',
      theme: new ThemeData(primaryColor: Colors.blue),
      home: Scaffold(
        appBar: AppBar(
          title: Text("Questions"),
        ),
        body: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                TextFormField(
                  textAlign: TextAlign.center,
                  style: TextStyle(color: Colors.blue, fontSize: 20),
                  controller: nameController,
                  validator: (value) {
                    if (value.isEmpty) {
                      return "Enter your name";
                    }
                  },
                  decoration: InputDecoration(
                    labelText: "Name",
                  ),
                ),
                TextFormField(
                  textAlign: TextAlign.center,
                  style: TextStyle(color: Colors.blue, fontSize: 20),
                  controller: roleController,
                  validator: (value) {
                    if (value.isEmpty) {
                      return "Enter your role";
                    }
                  },
                  decoration: InputDecoration(
                    labelText: "Cargo",
                  ),
                ),
                TextFormField(
                  keyboardType: TextInputType.number,
                  textAlign: TextAlign.center,
                  style: TextStyle(color: Colors.blue, fontSize: 20),
                  controller: salaryController,
                  validator: (value) {
                    if (value.isEmpty) {
                      return "Enter your salary";
                    }
                  },
                  decoration: InputDecoration(
                    labelText: "Salary",
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: salaryReajusted > 0
                      ? Text(
                          "Your adjusted salary is $salaryReajusted",
                          style: TextStyle(color: Colors.blue, fontSize: 20),
                        )
                      : null,
                ),
              ],
            ),
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: _calculateReajustSalary,
          tooltip: 'Calculate',
          child: Icon(Icons.send),
        ),
      ),
    );
  }

  Future<void> _calculateReajustSalary() async {
    var response = await QuestionsService.calculateReajustSalary(
      name: nameController.text,
      role: roleController.text,
      salary: double.parse(salaryController.text),
    );
    setState(() {
      salaryReajusted = double.parse(response.salary.toStringAsPrecision(2));
    });
  }
}

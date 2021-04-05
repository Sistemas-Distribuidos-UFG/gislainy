package server;

import idl.QuestionsIDL;


public class Questions  implements QuestionsIDL {

    public double CalculateReajustSalary(String name, String role, double salary) {
        if ("Operator".equals(role)) {
            salary *= 1.2;
        } else if ("Developer".equals(role)) {
            salary *= 1.2;
        }
        return salary;
    }
}

package client;

import idl.QuestionsIDL;

import java.rmi.NotBoundException;
import java.rmi.RemoteException;
import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;

public class Client {
    public static void main(String[] args) {
        try {
            Registry registry = LocateRegistry.getRegistry(55553);
            QuestionsIDL stub = (QuestionsIDL) registry.lookup("RMIQuestions");

            double salaryReajusted1 = stub.CalculateReajustSalary("Gislainy", "Developer", 5000);
            double salaryReajusted2 = stub.CalculateReajustSalary("Pabllo", "Operator", 3000);

            System.out.printf("Gislainy adjusted salary is %.2f\n", salaryReajusted1);
            System.out.printf("Pabllo adjusted salary is %.2f\n", salaryReajusted2);

        } catch (RemoteException | NotBoundException e) {
            e.printStackTrace();
        }

    }
}

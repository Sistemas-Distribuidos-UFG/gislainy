package idl;

import java.rmi.Remote;
import java.rmi.RemoteException;

public interface QuestionsIDL extends Remote {
    public double CalculateReajustSalary(String name, String role, double salary) throws RemoteException;

}

package server;

import idl.QuestionsIDL;

import java.rmi.RemoteException;
import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;
import java.rmi.server.UnicastRemoteObject;

public class Server  {
    public static void main(String[] args) {
        try {
            Questions questions = new Questions();

            QuestionsIDL stub = (QuestionsIDL) UnicastRemoteObject.exportObject(questions, 0);

            Registry registry = LocateRegistry.createRegistry(55553);

            registry.rebind("RMIQuestions", stub);
        } catch (RemoteException e) {
            e.printStackTrace();
        }
    }
}

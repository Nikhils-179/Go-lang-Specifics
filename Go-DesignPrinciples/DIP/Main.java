
public class Main {
    public static void main(String[] args) {
        ILogger fileLogger = new FileLogger();
        DataAccessLayer dal1 = new DataAccessLayer(fileLogger);
        dal1.addCustomer("John Doe"); 

        ILogger databaseLogger = new DatabaseLogger();
        DataAccessLayer dal2 = new DataAccessLayer(databaseLogger);
        dal2.addCustomer("Jane Doe");
    }
}

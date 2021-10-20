using System;

public static class Session
{
    public static string LoginUrl;
    public static string ResgisterUrl;
    public static string WebSocketUrl;
    public static string TcpUrl;
    public static string Token;

    public static string Account;

    public static UserData User=new UserData();
}


public class UserData
{
    public long UID;
    public string Icon;
    public string NickName;
    public UInt32 Gold;
    public UInt32 Diamond;
    public Pb.Enum.Sex Sex;
    
}
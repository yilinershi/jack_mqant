

using System;
using UnityEngine;
using UnityEngine.UI;

public class LoginView : MonoBehaviour
{
    public LoginView Instance;
    
    private Button btnLogin;
    private Button btnRegister;
    private InputField inputAccount;
    private InputField inputPassword;
    public void Start()
    {
        Instance = this;
        Debug.Log("login view start");
        btnLogin = transform.Find("Canvas/BtnLogin").GetComponent<Button>();
        btnRegister = transform.Find("Canvas/BtnRegister").GetComponent<Button>();
        inputAccount = transform.Find("Canvas/InputAccount").GetComponent<InputField>();
        inputPassword = transform.Find("Canvas/InputPassword").GetComponent<InputField>();
        btnLogin.onClick.AddListener(OnBtnLoginClick);
        btnRegister.onClick.AddListener(OnBtnRegisterClick);
    }

    private void OnBtnLoginClick()
    {
        LoginController.Login(inputAccount.text,inputPassword.text, () =>
        {
            //打开lobby
            var prefab = Resources.Load("Prefab/UILobby");
            var go = Instantiate(prefab) as GameObject;
            go.AddComponent<LobbyView>();
            
            //打开新的UI后关闭自己
            Destroy(this.gameObject);
        });
    }
    
    private void OnBtnRegisterClick()
    {
        LoginController.Register(inputAccount.text,inputPassword.text);
    }
    
}
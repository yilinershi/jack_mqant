using System;
using UnityEngine;
using UnityEngine.UI;
public class LobbyView : MonoBehaviour
{

    private Text LabelNickName;
    private Text LabelUserId;
    private Image ImageUerIcon;


    public void Start()
    {
        LabelNickName= transform.Find("Canvas/LabelNickName").GetComponent<Text>();
        LabelUserId= transform.Find("Canvas/LabelUserId").GetComponent<Text>();
        ImageUerIcon= transform.Find("Canvas/Image/UserIcon").GetComponent<Image>();
    }


    private void RefreshNickName()
    {
        
    }
}
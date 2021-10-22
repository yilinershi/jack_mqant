using System;
using UnityEngine;
using UnityEngine.UI;

public class LobbyView : MonoBehaviour
{
    private Text LabelNickName;
    private Text LabelUserId;
    private Image ImageUerIcon;
    private Text LabelDiamond;
    private Text LabelGold;

    private Button btnTetris;

    public void Start()
    {
        LabelNickName = transform.Find("Canvas/LabelNickName").GetComponent<Text>();
        LabelUserId = transform.Find("Canvas/LabelUserId").GetComponent<Text>();
        LabelDiamond = transform.Find("Canvas/LabelDiamond").GetComponent<Text>();
        LabelGold = transform.Find("Canvas/LabelGold").GetComponent<Text>();
        ImageUerIcon = transform.Find("Canvas/Image/UserIcon").GetComponent<Image>();
        btnTetris = transform.Find("Canvas/ButtonCreateTetris").GetComponent<Button>();
        RefreshAll();
    }

    private void RefreshAll()
    {
        RefreshDiamond();
        RefreshGold();
        RefreshNickName();
        RefreshUserId();
        btnTetris.onClick.AddListener(onBtnTetrisClick);
    }

    private void RefreshNickName()
    {
        LabelNickName.text = Session.User.NickName;
    }

    private void RefreshUserId()
    {
        LabelUserId.text = Session.User.UID.ToString();
    }

    private void RefreshDiamond()
    {
        LabelDiamond.text = "钻石:" + Session.User.Diamond;
    }

    private void RefreshGold()
    {
        LabelGold.text = "金币:" + Session.User.Gold;
    }

    private void onBtnTetrisClick()
    {
        TetrisController.CallSubscribeRoomInfo(true);
    }
}
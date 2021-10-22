using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class TetrisView : MonoBehaviour
{
    public static TetrisView Instance;

    private Button btnCreate;
    private Button btnJoin;
    private GameObject tableInfoGo;
    private Transform scrollRoot;
    private InputField labelTableName;

    public void Start()
    {
        Instance = this;
        btnCreate = transform.Find("Canvas/BtnCreate").GetComponent<Button>();
        btnJoin = transform.Find("Canvas/BtnJoin").GetComponent<Button>();
        btnCreate.onClick.AddListener(OnBtnCreateTetrisClick);
        btnJoin.onClick.AddListener(OnBtnJoinTetrisClick);
        tableInfoGo = transform.Find("Canvas/TableInfo").gameObject;
        scrollRoot = transform.Find("Canvas/Scroll View/Viewport/Content");
        labelTableName = transform.Find("Canvas/InputField").GetComponent<InputField>();
        RefreshTableList();
    }

    private void OnBtnCreateTetrisClick()
    {
        TetrisController.CallCreateTetris(labelTableName.text);
    }

    private void OnBtnJoinTetrisClick()
    {
        Debug.Log("OnBtnJoinTetrisClick");
    }

    public void RefreshTableList()
    {
        //清掉节点
        List<GameObject> list = new List<GameObject>();
        for (int i = 0; i < scrollRoot.childCount; i++)
        {
            list.Add(scrollRoot.GetChild(i).gameObject);
        }

        foreach (var child in list)
        {
            Destroy(child.gameObject);
        }


        //重新加上新的
        foreach (var tableInfo in TetrisController.allTableInfos)
        {
            var go = Instantiate(tableInfoGo, scrollRoot, true);
            go.SetActive(true);
            var nameLabel = go.transform.Find("Name").GetComponent<Text>();
            var creatorLabel = go.transform.Find("Creator").GetComponent<Text>();
            nameLabel.text = tableInfo.Name;
            creatorLabel.text = tableInfo.CreatorNickName;
        }
    }
}
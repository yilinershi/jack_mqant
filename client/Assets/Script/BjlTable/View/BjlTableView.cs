using System;
using System.Collections.Generic;
using UnityEngine;

public class TetrisTableView : MonoBehaviour
{

    private GameObject playerGo;
    private Transform playerRoot;

    private void Awake()
    {
        playerGo = transform.Find("Canvas/Player").gameObject;
        playerRoot = transform.Find("Canvas/PlayerRoot");
    }

    public void Refresh()
    {
        //清掉节点
        List<GameObject> list = new List<GameObject>();
        for (int i = 0; i < playerRoot.childCount; i++)
        {
            list.Add(playerRoot.GetChild(i).gameObject);
        }

        foreach (var child in list)
        {
            Destroy(child.gameObject);
        }
        
        foreach (var p in TetrisTableModel.AllPlayer)
        {
            var go = Instantiate(playerGo, playerRoot, true);
            go.SetActive(true);
            var playerView=    go.AddComponent<TetrisPlayerView>();
            playerView.Init(p);
        }
    }
}
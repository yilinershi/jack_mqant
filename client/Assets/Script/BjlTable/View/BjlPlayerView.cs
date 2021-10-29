using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class TetrisPlayerView : MonoBehaviour
{
    private Pb.Bjl.BjlPlayer data;
    private Text labelNickName;
    private Text labelScore;
    private Text labelLevel;

    public void Init(Pb.Bjl.BjlPlayer pushData)
    {
        data = pushData;
        labelNickName = transform.Find("NickName").GetComponent<Text>();
        labelScore = transform.Find("Score").GetComponent<Text>();
        labelLevel = transform.Find("Level").GetComponent<Text>();
        RefreshScore();
        RefreshNickName();
        RefreshLevel();
    }

    private void RefreshNickName()
    {
        labelNickName.text = "玩家：" + data.NickName;
    }

    private void RefreshScore()
    {
        labelScore.text = "分数：" + data.Score;
    }

    private void RefreshLevel()
    {
        labelLevel.text = "等级：" + data.Score;
    }
}
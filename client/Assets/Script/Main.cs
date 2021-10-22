using UnityEngine;
public class Main : MonoBehaviour
{
    private void Awake()
    {
        HttpNet.Handshake(() =>
        {
            var prefab = Resources.Load("Prefab/UILogin");
            var go = Instantiate(prefab) as GameObject;
            go.AddComponent<LoginView>();
            
        });
    }
}
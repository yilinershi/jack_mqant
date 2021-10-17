

using UnityEditor.VersionControl;
using UnityEngine;

public static class UIUtil
{
    public static void OpenUI<T>() where T :UIBase, new()
    {
        var name = "Panel_"+typeof(T);
        var prefab=  Resources.Load(name);

        var go = Object.Instantiate(prefab) as GameObject;
        var ui= go.AddComponent<T>();
        
    }
}
using System;
using System.Threading.Tasks;

public static class TaskUtil
{
   
    /// <summary>
    /// 构建Send的Task
    /// </summary>
    /// <param name="cb">回调</param>
    /// <param name="errorMsg">错误消息</param>
    /// <typeparam name="T">Proto类型</typeparam>
    /// <returns></returns>
    public static Task<T> GenSendTask<T>(out Action<T> cb, string errorMsg = "Pb error")
    {
        var tsc = new TaskCompletionSource<T>();
        cb = (T pb) =>
        {
            if (pb == null)
            {
                tsc.SetException(new Exception(errorMsg));
            }
            else
            {
                tsc.SetResult(pb);
            }
        };

        return tsc.Task;
    }

    
}
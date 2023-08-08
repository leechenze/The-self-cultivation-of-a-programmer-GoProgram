博学之, 审问之, 慎思之, 明辨之, 笃行之;
零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;







零.Kubernetes简介
    
    Github地址:
        https://github.com/gzyunke/test-k8s/tree/main
    博客地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/9EX8Cp45
    视频地址：
        https://www.bilibili.com/video/BV1Tg411P7EB/?vd_source=6122dea75af0b44c85ff18d196f1b32d
    视频地址：
        https://www.bilibili.com/video/BV1XY4y1t76G?p=83&vd_source=6122dea75af0b44c85ff18d196f1b32d


壹.安装 Kubernetes 集群

    章节参考的博客地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/nd7yOvdY
    CSDN裸机搭建K8s集群博客地址：
        https://blog.csdn.net/qq_41538097/article/details/124869179
    minikube官网：
        https://minikube.sigs.k8s.io/docs/commands/
    腾讯云官网：
        https://console.cloud.tencent.com/


贰.部署应用到集群中

    kubectl启动nginx应用（云平台搭建）（直接命令行运行的方式）
        部署命令：
            kubectl run nginx --image=nginx:latest
            kubectl get pod
            kubectl get pod -o wide
                查看pod包括对应的IP地址
        转发一下，让外部可以访问：
            kubectl port-forward --address 0.0.0.0 pod/nginx 80:80
                将本机80端口转发至pod的80端口
            curl 127.0.0.1:80 或 curl 127.0.0.1
                查看是否可以成功访问默认端口为80的nginx服务
            腾讯云平台点击当前的容器实例，选择 安全组/规则预览/入站规则/编辑规则/添加规则，配置开放80端口的防火墙，这样就可以使用公网进行访问了。
    
    章节参考地址：（配置文件运行的方式）
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/puf7fjYr


叁.Service
    
    章节参考地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/C0fakgwO


肆.StatefulSet
    
    章节参考地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/mJvk9q5z


伍.数据持久化(Storage)

    章节参考地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/Q2gBbjyW
    缺失命令记录：
        重启mongodb的statefulset的pod
        kubectl rollout restart statefulset mongodb
    sc, pvc, pv, statefulset配置都分开写了，写在一起也可以，在yaml中用---分割即可。


陆.ConfigMap&Secret(配置文件和密码)

    章节参考地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/YJf8LMtE
    缺失命令记录：
        mongo --host mongdb-0.mongodb -u mongouser -p mongopass

柒.Helm&命名空间
    
    Helm官网：
        https://helm.sh/zh/
    Helm应用中心：
        https://artifacthub.io/
    章节参考地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/3iQiyInr


捌.Ingress
    
    章节参考地址：
        https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/JAxVnLJ5



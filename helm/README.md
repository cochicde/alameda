# What is Alameda

Alameda is a prediction engine that foresees future resource usage of your Kubernetes cluster from the cloud layer down to the pod level. We use machine learning technology to provide intelligence that enables dynamic scaling and scheduling of your containers - effectively making us the “brain” of Kubernetes resource orchestration. By providing full foresight of resource availability, demand, health, impact and SLA, we enable cloud strategies that involve changing provisioned resources in real time. 

For more details, please refer to https://github.com/containers-ai/alameda

# Alameda deployment with Helm chart

> **Note**: To deploy Alameda by Helm charts, please install [Helm](https://docs.helm.sh/using_helm/#quickstart-guide) first.

According to Alameda [architecture](https://github.com/containers-ai/alameda/blob/master/design/architecture.md), it is composed of:
- alameda-operator
- alameda-datahub
- alameda-ai
- alameda-influxdb (leverage the opensource InfluxDB)

and assumes **Prometheus** is running in your cluster.

Users can install Alameda by following:
1. Install InfluxDB chart by executing:
```
$ helm install stable/influxdb --version 1.3.3 --set persistence.enabled=false --name alameda-influxdb --namespace alameda
```
2. Install cert-manager
Please refer to [install guide](https://github.com/jetstack/cert-manager/blob/release-0.9/deploy/charts/cert-manager/README.md) for installing cert-manager.
```
$ kubectl apply \
    -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.9/deploy/manifests/00-crds.yaml
$ kubectl label namespace cert-manager certmanager.k8s.io/disable-validation="true"
$ helm repo add jetstack https://charts.jetstack.io
$ helm install --name cert-manager --namespace cert-manager jetstack/cert-manager
```
3. Install Alameda chart for component _alameda-operator_, _alameda-datahub_, _alameda-ai_, and _alameda-ai-dispatcher_ _alameda_notifier by executing:
```
$ helm install --name alameda --namespace alameda ./alameda
```
> **Note 1**: Alameda needs to collaborate with Prometheus to see historical pod/node metrics. The default URL is set to _http://prometheus-prometheus-oper-prometheus.monitoring:9090_ in this chart. Please modify it according to your environment before installing this chart.  
> **Note 2**: The images of Alameda components are assumed existed in local docker environment and pulled from it. Please refer to [build guide](../docs/build.md) for building images from source code or change the image repository settings before installing this chart.

4. (Optional) If your environment does not have a running Prometheus, you can install it by executing:
```
$ helm install stable/prometheus-operator --version 6.4.4 --name prometheus --namespace monitoring
```
This will install Prometheus and the default setting will have all the metrics that Alameda needs. For detail metrics needed by Alameda, please visit [metrics_used_in_Alameda.md](../docs/metrics_used_in_Alameda.md) document.
> **Note**: prometheus chart version greater than 6.4.4 has rules changed which make datahub work abnormal in kubernetes version 1.14 or later.

## TL;DR;

```console
$ git clone https://github.com/containers-ai/alameda
$ helm install stable/influxdb --version 1.3.3 --set persistence.enabled=false --name alameda-influxdb --namespace alameda
### Install Prometheus if no existed one
# $ helm install stable/prometheus-operator --version 6.4.4 --name prometheus --namespace monitoring
$ helm install --name alameda --namespace alameda ./alameda/
```


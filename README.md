# Adaptec Hardware RAID Controller Prometheus Exporter

### Overview

This is a Prometheus exporter for Adaptec hardware RAID controllers, written in Go. The exporter provides a variety of metrics that offer insight into the health and performance of your RAID arrays, ready to be scraped by Prometheus.

### Metrics

| Metric                                 | Description                         |
| -------------------------------------- | ----------------------------------- |
| `adaptec_logical_device_status`        | Status of Logical Devices           |
| `adaptec_physical_device_state`        | State of Physical Devices           |
| `adaptec_physical_device_smart_warnings`| SMART warnings of Physical Devices |

### Prerequisites

- Adaptec Hardware RAID controller
- Go 1.18+
- Prometheus installed in your environment
- Adaptec `arcconf` command line utility

### Installation

```bash
go install github.com/d9ff/adaptec-raid-exporter
```

### Usage

Launch the exporter using the following command:

```bash
$GOPAT/bin/adaptec-raid-exporter --web.listen-address=0.0.0.0:9101
```

The exporter is set to run on port `9101` by default, but you can modify the `--web.listen-address` flag to use a different port.

To check that the exporter is running and serving metrics, curl the `/metrics` endpoint:

```bash
curl localhost:9101/metrics
```

### Prometheus Configuration

To scrape metrics from this exporter, add the following job to your `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'adaptec_raid'
    static_configs:
      - targets: ['<ip-of-your-server>:9101']
```

Replace `<ip-of-your-server>` with the IP address or DNS name of your RAID controller.

Then restart Prometheus:

```bash
systemctl restart prometheus
```

### Alerting

An example alerting rules file for Prometheus is provided in `alerts/adaptec_raid.yml`. You can modify this file as per your needs.

### Contributing

We welcome contributions to this project. Feel free to open issues or pull requests for bugs, features, or any other changes.

### License

This project is licensed under the MIT License.

### Contact

For any questions or concerns, open an issue on this GitHub repository.

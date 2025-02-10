# [Google Cloud Monitoring](https://cloud.google.com/monitoring) [Notification Channel](https://cloud.google.com/monitoring/alerts/using-channels-api) for [Pushover](https://pushover.net))

[![build-container](https://github.com/DazWilkin/go-gcp-pushover-notificationchannel/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/go-gcp-pushover-notificationchannel/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/go-gcp-pushover-notificationchannel.svg)](https://pkg.go.dev/github.com/DazWilkin/go-gcp-pushover-notiificationchannel)
[![Go Report Card](https://goreportcard.com/badge/github.com/DazWilkin/go-gcp-pushover-notificationchannel)](https://goreportcard.com/report/github.com/DazWilkin/go-gcp-pushover-notificationchannel)
+ `ghcr.io/dazwilkin/go-gcp-pushover-notificationchannel:41364588831b35fd10b8971a384deb96051449c1`

See [Using Google Monitoring Alerting to send Pushover notifications](https://pretired.dazwilkin.com/posts/220514/)

## Prerequisites

### Required Secrets
Before deploying, you need to set up the following secrets in Google Cloud Secret Manager:

1. **Pushover Token** (`PUSHOVER_A`)
   - Your Pushover application token
   - Used to authenticate with Pushover API
   - Create with:
     ```bash
     echo -n "your-pushover-app-token" | gcloud secrets create PUSHOVER_A --data-file=-
     ```

2. **Pushover User ID** (`PUSHOVER_USER_ID`)
   - Your Pushover user key
   - Used to identify which user/device to send notifications to
   - Create with:
     ```bash
     echo -n "your-pushover-user-key" | gcloud secrets create PUSHOVER_USER_ID --data-file=-
     ```

These secrets are automatically mapped to the following environment variables in the application:
- `PUSHOVER_TOKEN`: Maps to the `PUSHOVER_A` secret
- `PUSHOVER_USER_ID`: Maps to the `PUSHOVER_USER_ID` secret

To get your Pushover credentials:
1. Log in to your [Pushover account](https://pushover.net)
2. Create a new application to get the application token
3. Your user key is shown on the main page after logging in

## Setup Instructions

### 1. Deploy the Service
The service should be deployed to Cloud Run. You can use Cloud Build or your preferred deployment method.

### 2. Configure Notification Channel
1. **Get the Service URL**
   - Go to Google Cloud Console > Cloud Run
   - Find your `notificationchannel` service
   - Copy the service URL (format: `https://notificationchannel-xxxxx-xx.a.run.app`)

2. **Create a Notification Channel**
   - Navigate to Google Cloud Console > Monitoring > Alerting > Notification channels
   - Click "Add New"
   - Select "Webhook"
   - Configure the following:
     - Name: `Pushover Notifications`
     - Endpoint URL: [Your Cloud Run service URL]
     - HTTP Headers: `Content-Type: application/json`

### 3. Create Alert Policies
1. Go to Monitoring > Alerting
2. Click "Create Policy"
3. Configure your desired monitoring conditions
4. In the Notifications section:
   - Add your newly created webhook notification channel
5. Save the policy

### 4. Testing
- You can test the integration by:
  - Setting a temporary sensitive threshold in your alert policy
  - When triggered, you should receive Pushover notifications on your configured devices
- The notification will include:
  - Formatted alert data
  - Relevant links and details
  - Incident information

## [Sigstore](https://sigstore.dev)

`go-gcp-pushover-notificationchannel` container images are being signed by Sigstore and may be verified:

```bash
cosign verify \
--key=./cosign.pub \
ghcr.io/dazwilkin/go-gcp-pushover-notificationchannel:41364588831b35fd10b8971a384deb96051449c1
```

> **NOTE** [`cosign.pub`](/cosign.pub) may be downloaded here

To install cosign, e.g.:

```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```


<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

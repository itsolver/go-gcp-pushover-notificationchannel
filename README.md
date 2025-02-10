# [Google Cloud Monitoring](https://cloud.google.com/monitoring) [Notification Channel](https://cloud.google.com/monitoring/alerts/using-channels-api) for [Pushover](https://pushover.net))

[![build-container](https://github.com/DazWilkin/go-gcp-pushover-notificationchannel/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/go-gcp-pushover-notificationchannel/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/go-gcp-pushover-notificationchannel.svg)](https://pkg.go.dev/github.com/DazWilkin/go-gcp-pushover-notiificationchannel)
[![Go Report Card](https://goreportcard.com/badge/github.com/DazWilkin/go-gcp-pushover-notificationchannel)](https://goreportcard.com/report/github.com/DazWilkin/go-gcp-pushover-notificationchannel)

A webhook notification channel for Google Cloud Monitoring that forwards alerts to Pushover. When alerts trigger in Cloud Monitoring, they will be sent as push notifications to your devices via Pushover.

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

To get your Pushover credentials:
1. Log in to your [Pushover account](https://pushover.net)
2. Create a new application to get the application token
3. Your user key is shown on the main page after logging in

## Setup Instructions

### 1. Deploy the Service
The service is automatically deployed to Cloud Run using Cloud Build. The included `cloudbuild.yaml` will:
1. Build the container image
2. Push it to Artifact Registry
3. Deploy to Cloud Run with the required configuration

To deploy:
```bash
git clone https://github.com/itsolver/go-gcp-pushover-notificationchannel.git
cd go-gcp-pushover-notificationchannel
git push  # This will trigger Cloud Build to deploy
```

### 2. Configure Notification Channel
1. **Get the Service URL**
   - Go to Google Cloud Console > Cloud Run
   - Find your `notificationchannel` service
   - Copy the service URL (format: `https://notificationchannel-xxxxx-xx.a.run.app`)
   - **Important**: Add `/webhook` to the end of the URL

2. **Create a Notification Channel**
   - Navigate to Google Cloud Console > Monitoring > Alerting > Notification channels
   - Click "Add New"
   - Select "Webhook"
   - Configure the following:
     - Display Name: `Pushover Notifications` (or your preferred name)
     - Endpoint URL: `https://your-service-url/webhook`
     - Leave "Use HTTP Basic Auth" unchecked
   - Click "Test Connection" to verify setup

### 3. Create Alert Policies
1. Go to Monitoring > Alerting
2. Click "Create Policy"
3. Configure your desired monitoring conditions
4. In the Notifications section:
   - Add your newly created webhook notification channel
5. Save the policy

### 4. Testing
The webhook can be tested in two ways:
1. **Using the Test Connection button**
   - In the webhook configuration
   - Should receive an immediate test notification on your Pushover devices
2. **Using an Alert Policy**
   - Create a policy with a sensitive threshold
   - When the condition triggers, you'll receive formatted notifications including:
     - Alert state (open/closed)
     - Project ID
     - Summary of the incident
     - Timing information
     - System labels and metadata

## Troubleshooting
If you're not receiving notifications:
1. Check Cloud Run logs for any errors
2. Verify the webhook URL includes the `/webhook` path
3. Confirm your Pushover credentials are correctly set in Secret Manager
4. Test the webhook connection from the notification channel configuration

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

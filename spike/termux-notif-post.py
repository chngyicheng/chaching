import json
import subprocess

result = subprocess.run(["termux-notification-list"], capture_output=True, text=True)

if result.returncode == 0:
    notifications = json.loads(result.stdout)
    for n in notifications:
        print(f"Text: {n.get('content')}")
else:
    print("Failed to get notifications: ", result.stderr)

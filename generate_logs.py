import os
import random
from datetime import datetime, timedelta

def generate_log_file(file_path, server_id, num_lines=500):
    start_time = datetime(2025, 7, 15, 0, 0, 0)
    log_levels = ["INFO", "DEBUG", "ERROR", "WARN"]
    users = ["alice", "bob", "carol", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy", "mallory", "niaj", "oscar", "peggy", "rupert", "sybil", "trent", "victor", "walter", "yvonne", "zach"]

    with open(file_path, "w") as f:
        for i in range(num_lines):
            timestamp = start_time + timedelta(seconds=i * 10)
            log_level = random.choice(log_levels)
            user = random.choice(users)
            action = random.choice([
                f"User {user} logged in",
                f"User {user} logged out",
                f"User {user} updated profile",
                f"User {user} changed password",
                f"Server {server_id} started",
                f"Server {server_id} shutdown",
                f"Disk error on /dev/sda{server_id}",
                f"High CPU usage detected",
                f"Backup started",
                f"Backup completed",
                f"Query executed: SELECT * FROM logs WHERE user='{user}'",
                f"User {user} uploaded file report.pdf",
                f"User {user} deleted file temp.tmp",
                f"Server {server_id} received request from 10.0.0.{random.randint(1, 255)}",
                f"Server {server_id} heartbeat OK",
                f"User {user} accessed dashboard",
                f"User {user} failed login attempt",
                f"Server {server_id} encountered memory leak",
                f"Server {server_id} restarted",
                f"User {user} initiated data export",
                f"User {user} imported data.csv",
                f"Server {server_id} completed log rotation",
                f"User {user} triggered security alert",
                f"Server {server_id} updated configuration",
                f"User {user} performed system health check",
            ])
            f.write(f"[{log_level}] {action} at {timestamp.strftime('%Y-%m-%d %H:%M:%S')}\n")

def main():
    log_dir = "log_files"
    os.makedirs(log_dir, exist_ok=True)

    for server_id in range(1, 11):
        file_path = os.path.join(log_dir, f"machine.{server_id}.log")
        generate_log_file(file_path, server_id)

if __name__ == "__main__":
    main()

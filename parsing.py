import subprocess

def open_parsehub():
    # Path to the ParseHub application on Mac
    app_path = "/Applications/ParseHub.app"
    subprocess.run(["open", app_path])

open_parsehub()

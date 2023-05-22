use std::{
    collections::HashMap,
    env,
    io::{self, ErrorKind},
    process::{Command, Output},
};

use anyhow::{bail, Result};
use furse::{structures::ID, Furse};
use url::Url;

const CURSEFORGE_API_KEY: &str = env!("CURSEFORGE_API_KEY");

#[tokio::main]
async fn main() -> Result<()> {
    #[cfg(target_os = "windows")]
    try_update_registry()?;

    let mut args: Vec<String> = env::args().collect();
    args.drain(0..1);
    if args.is_empty() {
        bail!("Missing arguments")
    }

    let url = Url::parse(&args.join(" "))?;
    let query = url.query_pairs().collect::<HashMap<_, _>>();
    let Some(addon_id) = query.get("addonId").and_then(|id| id.parse::<ID>().ok()) else {
        bail!("Missing or Malformed query parameter: addonId")
    };
    let Some(mod_id) = query.get("fileId").and_then(|id| id.parse::<ID>().ok()) else {
        bail!("Missing or Malformed query parameter: fileId")
    };

    let api_key = env::var("CURSEFORGE_API_KEY").unwrap_or_else(|_| CURSEFORGE_API_KEY.into());
    let curseforge = Furse::new(&api_key);
    let download_url = curseforge.file_download_url(addon_id, mod_id).await?;

    // You may declare your MultiMC based launcher variant here
    // Please make sure the capitalization matches the filename
    // Windows and Linux are case sensitive, macOS is not
    #[allow(unused_mut)]
    let mut launchers = vec!["MultiMC", "multimc", "polymc", "prismlauncher"];

    // Why does Petr Mrázek (Peterix) die on the stupidest hills
    #[cfg(target_os = "linux")]
    launchers.insert(0, "/opt/multimc/run.sh");

    #[cfg(target_os = "linux")]
    match try_flatpaks(download_url.as_ref()) {
        Ok(true) => return Ok(()),
        Ok(false) => {}
        Err(e) => {
            // NotFound indicates flatpak is not installed
            if ErrorKind::NotFound != e.kind() {bail!(e)}
        }
    };

    for launcher in launchers {
        match try_launcher(launcher, download_url.as_ref()) {
            Ok(_) => return Ok(()),
            Err(error) => {
                if let ErrorKind::NotFound = error.kind() {
                    continue;
                } else {
                    bail!(error)
                }
            }
        }
    }

    bail!("Failed to find launcher")
}

#[cfg(target_os = "windows")]
fn try_update_registry() -> Result<()> {
    use std::path::Path;

    use is_elevated::is_elevated;
    use winreg::{enums::HKEY_CLASSES_ROOT, RegKey};

    if is_elevated() {
        let exe = env::current_exe()?;
        let value = format!("\"{}\" \"%1\"", exe.to_string_lossy());

        let hkcr = RegKey::predef(HKEY_CLASSES_ROOT);

        let root_path = Path::new("curseforge");
        let (root_key, _) = hkcr.create_subkey(root_path)?;
        root_key.set_value("URL Protocol", &"")?;

        let sub_path = root_path.join("shell\\open\\command");
        let (sub_key, _) = hkcr.create_subkey(sub_path)?;
        sub_key.set_value("", &value)?;

        println!("Registry Updated");
    }

    Ok(())
}

#[cfg(target_os = "linux")]
fn try_flatpaks(download_url: &str) -> Result<bool, io::Error> {
    let packages = vec!["org.polymc.PolyMC", "org.prismlauncher.PrismLauncher"];
    for package in packages {
        let output = Command::new("flatpak")
            .args(["run", package, "--import", download_url])
            .output()?;
        let Some(code) = output.status.code() else {
            continue;
        };
        if code == 0 {
            return Ok(true);
        }
    }

    Ok(false)
}

#[cfg(target_os = "windows")]
fn try_launcher(launcher: &str, download_url: &str) -> Result<Output, io::Error> {
    Command::new(launcher.to_owned() + ".exe")
        .args(["--import", download_url])
        .output()
}

#[cfg(target_os = "macos")]
fn try_launcher(launcher: &str, download_url: &str) -> Result<Output, io::Error> {
    Command::new("open")
        .args(["-a", launcher, "--args", "--import", download_url])
        .output()
}

#[cfg(target_os = "linux")]
fn try_launcher(launcher: &str, download_url: &str) -> Result<Output, io::Error> {
    Command::new(launcher)
        .args(["--import", download_url])
        .output()
}

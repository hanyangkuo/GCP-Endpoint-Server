copy "..\bin\osquery-4.7.0.msi" ".\"
msiexec /i "osquery-4.7.0.msi" /l* "C:\Program Files\osquery\OsqueryInstall.log" /passive
del osquery-4.7.0.msi
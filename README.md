# moe_priv_server_fix

Workaround for private server not showing up in Myth of Empires server browser.

Works by responding to server list query that usually goes to http://l11-prod-list-hk-myth.dyhxgame.com/private_list. Instructions:

1. Add a host file entry in c:\Windows\System32\drivers\etc\hosts for the above domain, and point at local host. Example:
    # localhost name resolution is handled within DNS itself.
    #	127.0.0.1       localhost
    #	::1             localhost
    127.0.0.1 l11-prod-list-hk-myth.dyhxgame.com
2. Grab the .exe in releases section, or build from source - go build.
3. Edit server_list.json and change IP/Port to your server. Will likely need to update the Major / Minor revision each patch to match client.
4. Run .exe, it might prompt to add a firewall exception - you'll need to allow
5. Open Myth of empires - server browser will now only display the server(s) you've added to the list

To undo it, just delete the host entry in step 1. This will prevent MOE from using local machine server list.

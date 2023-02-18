<h1>neteye V2</h1>
<p>With neteye you can search in your network or others networks for a specific port.<p>
<p>Now it's in golang, now it's faster</p>
<p>It will scan a range of ip address searching that port</p>
<h2>Usage:</h2>
<p>You need to provide a valid network such as 192.168.1.0 or 172.16.0.0 the netmask is default /24 but you can change it adding -m 16 or 8</p>
<p>Only /8, /16, /24 /0 (Global) are supported by now</p>
<h4>Example: </h4>
<p>neteye -i 192.168.1.0 -p 80</p>
<h4>Example with netmask: </h4>
<p>neteye -i 172.16.0.0 -p 80 -m 16</p>
<h4>Threads</h4>
<p>Default threads -> 200 you can change it with the -t parameter</p>
<h3>Global mode</h3>
<p>This type of scan its about all the public ipv4 address</p>
<p>neteye -i 172.16.0.0 -p 80 -m 16 -t 80</p>
<h4>Example: </h4>
<p>neteye -p 22 -m 0</p>

<h2>Installation</h2>
<p>Just execute the script install.sh as root</p>

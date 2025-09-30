# RemoteShellClient
Aplicación del cliente de un sistema de administración remota de computadoras, la aplicación cliente permite a un computador con sistema operativo Windows ser administrado por otro.

#INSTRUCCIONES DE USO REMOTESHELL:

Colocar el paquete RemoteShellClient en un entorno Windows

Colocar el paquete RemoteShellServer en un entorno Unix

del paquete RemoteShellServer, ubicar los archivos:
RemoteShellServer(ejecutable) -> /usr/bin
RemoteShellServer.conf -> /etc
users.db -> /var/lib/RemoteShellServer

Editar el archivo /etc/RemoteShellServer.conf, poner la ip de su maquina y el puerto deseado


##PARA INCIAR EL PROGRAMA
Ejecutar Server en Unix (comando RemoteShellServer)
Ejecutar Cliente ejecutar ClientRemoteShell.exe seguido de los parametros <ip> <puerto> <tiempo>
*ip: Direccion ip de la máquina a la que se quiere conectar
*puerto: puerto por el que escucha el server
*tiempo: tiempo entre reportes de uso de recursos
ejemplo:
$./ClienteRemoteShell.exe 192.168.1.22 45345 10

import shutil
import os
ruta="C:\\Users\\cueto\\Downloads"
descargas=os.listdir(ruta)
#print(os.listdir(ruta))
carpeta = input("Ingrese el nombre de la carpeta:\n")
if not os.path.isdir(ruta+"\\"+ carpeta): 
    os.mkdir(ruta+"\\" + carpeta) # mkdir -> crea una carpeta en una ruta dada

newRuta=ruta+"\\" +carpeta

for i,archivo in enumerate(descargas):
    #print(i,archivo)
    print(f"{newRuta}\\{archivo}")
    print(archivo.endswith("jpeg"))
    print(archivo)
    if ((archivo.endswith('.docx')) ):
        print('A')
        # print(file_name)
        shutil.move(ruta + f'\\{archivo}', newRuta + f'\\{archivo}')
        print("done")
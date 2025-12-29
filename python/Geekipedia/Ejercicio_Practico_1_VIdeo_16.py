
print("-------------------Sistema de Vacaciones Rappi-----------------")


nombre = input("Ingrese su nombre: ")
clave_departamento = int(input("Ingrese la clave de su departamento: "))
antiguedad = int(input("Ingrese los años de antiguedad en la empresa: "))

def mostra_datos(nombre,antiguedad):
    print("Empleado: ",nombre)
    print("Años de antiguedad: ",antiguedad)


match clave_departamento:
    case 1:
        print("Perteneces al departamento de Atención al cliente")

        if antiguedad == 1:
            print("Te corresponden 6 dias de vacaciones",)
        elif antiguedad >= 2 and antiguedad <=6:
            print("Te corresponden 14 dias de vacaciones")
        elif antiguedad>=7:
            print("Te corresponden 20 dias de vacaciones")


    case 2:
        print("Pertenece al departamento de Logistica")
    case 3:
        print("Pertenece al departamento de la Gerencia")
    case _:
        print("No perteneces a ningun departamento")





"""


print("Nombre: ",nombre)
print("Clave del departamento al cual pertence: ",clave_departamento)
print("Años de antiguedad",antiguedad)

"""
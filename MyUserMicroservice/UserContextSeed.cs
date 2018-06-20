using Microsoft.AspNetCore.Builder;
using Microsoft.EntityFrameworkCore;
using MyUserMicroservice.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;

namespace MyUserMicroservice
{
    public class UserContextSeed
    {

       

       

        public static async Task SeedAsync(MyUserContext contextuser)
        {


            var context = contextuser;




            using (context)
            {
                context.Database.Migrate();
                if (!context.User.Any())
                {
                    context.User.AddRange(
                        GetPreconfiguredUsers());
                    await context.SaveChangesAsync();
                }
               
            }
        }

        static IEnumerable<User> GetPreconfiguredUsers()
        {
            return new List<User>()
       {
           new User() { Nombre="Usuario 1", Email="Email 1", Ciudad="Ciudad 1", Direccion="Direccion 1", NoTel="Tel 1", Pais="Pais 1", Password=MD5Hash("Password 1"), Verificado=true},
           new User() { Nombre="Usuario 2", Email="Email 2", Ciudad="Ciudad 2", Direccion="Direccion 2", NoTel="Tel 2", Pais="Pais 2", Password=MD5Hash("Password 2"), Verificado=false },
           new User() { Nombre="Usuario 3", Email="Email 3", Ciudad="Ciudad 3", Direccion="Direccion 3", NoTel="Tel 3", Pais="Pais 3", Password=MD5Hash("Password 3"), Verificado=true },
           new User() { Nombre="Usuario 4", Email="Email 4", Ciudad="Ciudad 4", Direccion="Direccion 4", NoTel="Tel 4", Pais="Pais 4", Password=MD5Hash("Password 4"), Verificado=false }
       };
        }

        public static string MD5Hash(string input)
        {
            using (var md5 = MD5.Create())
            {
                var result = md5.ComputeHash(Encoding.ASCII.GetBytes(input));
                return Encoding.ASCII.GetString(result);
            }
        }


    }
}

﻿using System;
using System.Collections.Generic;

namespace MyUserMicroservice.Models
{
    public partial class User
    {
        public int Id { get; set; }
        public string Nombre { get; set; }
        public string Email { get; set; }
        public string Password { get; set; }
        public bool? Verificado { get; set; }
        public string NoTel { get; set; }
        public string Pais { get; set; }
        public string Ciudad { get; set; }
        public string Direccion { get; set; }
    }
}

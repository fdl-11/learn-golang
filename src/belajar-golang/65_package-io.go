/*
IO (Input Output) merupakan fitur di golang yang digunakan sebagai standard untuk proses Input Output

[Reader]
Untuk membaca input, golang menggunakan kontrak interface bernama Reader yang terdapat di package io

[Writer]
Untuk menulis ke input, golang menggunakan kontrak interface bernama Writer yang terdapat di package io

[Implementasi IO]
Penggunaan dari IO terdapat di banyak package, sebelumnya sudah pernah digunakan di materi package encoding csv
Karena package IO sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri
Tapi untungnya, golang juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu menggunakan package bufio
*/

# Tugas-5_NetPro-kelompok-

  <a><strong>  ANGGOTA KELOMPOK :  </strong></a>
  <p><a>  Meilyand Evriyan Timor (1301161769)  </a> </p>
  <p><a>  Reyhan Rahmansyah (1301160805)  </a> </p>
  <p><a>  Reno Butar Butar  (1301164724)  </a> </p>


#### HASIL RUNNING NOMOR 1 ####

[![JSONMarshal.png](https://i.postimg.cc/Cxg6W1ZM/JSONMarshal.png)](https://postimg.cc/crFm3dMP)

Cara kerjanya :
Fungsi json.Marshal digunakan untuk decoding data ke json string. Sumber data bisa berupa variabel objek cetakan struct, map[string]interface{}, atau slice.
 Kemudian Struct Person diinisialisasi pada variabel bytes dengan atribut nama depan (FirstName) “John” dan nama belakang (LastName) “Dow”. Kemudian, bytes diserialisasikan kedalam bentuk JSON dan dicetak menghasilkan keluaran struct Person dalam bentuk JSON ({“FirstName”:”John”,”LastName”:”Dow”}).

#### HASIL RUNNING NOMOR 2 ####

[![JSONUNmarshal.png](https://i.postimg.cc/fbF7Q2H7/JSONUNmarshal.png)](https://postimg.cc/d72Z8BcD)

Cara Kerja  : 
Data json dituliskan sebagai string. Dengan menggunakan json.Unmarshal, json string bisa dikonversi menjadi bentuk objek. Struct Person dalam bentuk JSON di-assign ke variabel in, yang kemudian variabel in menjadi nilai untuk variabel bytes dalam tipe data []byte. Kemudian, variabel bytes di decode dari bentuk JSON menjadi bentuk struct Person dan ditampung pada variabel p. Sehinga, keluaran yang dihasilkan dari mencetak variabel p adalah {Firstname:John Lastname:Dow}.


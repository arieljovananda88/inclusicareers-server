CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS properti (
    "id_properti" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "nama_properti" varchar NOT NULL,
    "harga_dasar" BIGINT NOT NULL,
    "harga_2022" BIGINT NOT NULL,
    "harga_2021" BIGINT NOT NULL,
    "harga_2020" BIGINT NOT NULL,
    "harga_2019" BIGINT NOT NULL,
    "tipe" varchar NOT NULL,
    "area" INT NOT NULL,
    "kondisi" varchar NOT NULL,
    "alamat" varchar NOT NULL,
    "biaya_sewa" BIGINT NOT NULL,
    "pengali" FLOAT NOT NULL,
    "deskripsi_bisnis" varchar NOT NULL,
    "deskripsi_pribadi" varchar NOT NULL,
    "name_agen" varchar NOT NULL,
    "nomor_agen" varchar NOT NULL,
    "lat_position" FLOAT NOT NULL,
    "long_position" FLOAT NOT NULL,
    "link_map" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
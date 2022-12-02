PGDMP                         z            db_personal_web    15.1    15.1                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            	           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            
           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    24577    db_personal_web    DATABASE     �   CREATE DATABASE db_personal_web WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE db_personal_web;
                postgres    false            �            1259    24704    tb_projects    TABLE     6  CREATE TABLE public.tb_projects (
    id_project integer NOT NULL,
    name character varying NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    description text NOT NULL,
    technologies character varying[] NOT NULL,
    image character varying NOT NULL,
    id_user integer NOT NULL
);
    DROP TABLE public.tb_projects;
       public         heap    postgres    false            �            1259    24702    tb_projects_id_project_seq    SEQUENCE     �   CREATE SEQUENCE public.tb_projects_id_project_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.tb_projects_id_project_seq;
       public          postgres    false    218                       0    0    tb_projects_id_project_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.tb_projects_id_project_seq OWNED BY public.tb_projects.id_project;
          public          postgres    false    216            �            1259    24703    tb_projects_id_user_seq    SEQUENCE     �   CREATE SEQUENCE public.tb_projects_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.tb_projects_id_user_seq;
       public          postgres    false    218                       0    0    tb_projects_id_user_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.tb_projects_id_user_seq OWNED BY public.tb_projects.id_user;
          public          postgres    false    217            �            1259    24694    tb_users    TABLE     �   CREATE TABLE public.tb_users (
    id_user integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL
);
    DROP TABLE public.tb_users;
       public         heap    postgres    false            �            1259    24693    tb_users_id_user_seq    SEQUENCE     �   CREATE SEQUENCE public.tb_users_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.tb_users_id_user_seq;
       public          postgres    false    215                       0    0    tb_users_id_user_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.tb_users_id_user_seq OWNED BY public.tb_users.id_user;
          public          postgres    false    214            l           2604    24707    tb_projects id_project    DEFAULT     �   ALTER TABLE ONLY public.tb_projects ALTER COLUMN id_project SET DEFAULT nextval('public.tb_projects_id_project_seq'::regclass);
 E   ALTER TABLE public.tb_projects ALTER COLUMN id_project DROP DEFAULT;
       public          postgres    false    216    218    218            m           2604    24708    tb_projects id_user    DEFAULT     z   ALTER TABLE ONLY public.tb_projects ALTER COLUMN id_user SET DEFAULT nextval('public.tb_projects_id_user_seq'::regclass);
 B   ALTER TABLE public.tb_projects ALTER COLUMN id_user DROP DEFAULT;
       public          postgres    false    218    217    218            k           2604    24697    tb_users id_user    DEFAULT     t   ALTER TABLE ONLY public.tb_users ALTER COLUMN id_user SET DEFAULT nextval('public.tb_users_id_user_seq'::regclass);
 ?   ALTER TABLE public.tb_users ALTER COLUMN id_user DROP DEFAULT;
       public          postgres    false    214    215    215                      0    24704    tb_projects 
   TABLE DATA           x   COPY public.tb_projects (id_project, name, start_date, end_date, description, technologies, image, id_user) FROM stdin;
    public          postgres    false    218   "                 0    24694    tb_users 
   TABLE DATA           B   COPY public.tb_users (id_user, name, email, password) FROM stdin;
    public          postgres    false    215   (                  0    0    tb_projects_id_project_seq    SEQUENCE SET     H   SELECT pg_catalog.setval('public.tb_projects_id_project_seq', 4, true);
          public          postgres    false    216                       0    0    tb_projects_id_user_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.tb_projects_id_user_seq', 1, false);
          public          postgres    false    217                       0    0    tb_users_id_user_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.tb_users_id_user_seq', 2, true);
          public          postgres    false    214            q           2606    24712    tb_projects tb_projects_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public.tb_projects
    ADD CONSTRAINT tb_projects_pkey PRIMARY KEY (id_project);
 F   ALTER TABLE ONLY public.tb_projects DROP CONSTRAINT tb_projects_pkey;
       public            postgres    false    218            o           2606    24701    tb_users tb_users_pkey 
   CONSTRAINT     Y   ALTER TABLE ONLY public.tb_users
    ADD CONSTRAINT tb_users_pkey PRIMARY KEY (id_user);
 @   ALTER TABLE ONLY public.tb_users DROP CONSTRAINT tb_users_pkey;
       public            postgres    false    215            r           2606    24713 $   tb_projects tb_projects_id_user_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.tb_projects
    ADD CONSTRAINT tb_projects_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.tb_users(id_user);
 N   ALTER TABLE ONLY public.tb_projects DROP CONSTRAINT tb_projects_id_user_fkey;
       public          postgres    false    218    215    3183               �  x�mT�nG=���3IH�s�)�� 6`�m�R�.[ӛzqB���j��� `�S]o�Ws9��`����"�-�8Kv����rs�ny���7U������L�P���'�\��q�8sX'gg.��g��N֧^�4=3��Q+�	����}�N�j��p�l�G�PQ�q8�v��X��u=�r��/�,<���y#g�zV�Ma.W��������߼h޽���pÎ�0�)�G�iM�������ù�7�-]�
}6t��>��qReK�yV3��9q�h/��	с�v�������;#�[m��fϘG�=OlA�Xֻgc9l��TN�Z��UX��H�^�''{�z�s+:�9�.a��\�v�\��	��1�������s3��m߭��O,��xu��x���%�n$�_��
m8��;�[������d;�ÝIM��#��sW6�:m��'�W���{�J�yZ�����̽�s T��[2}�v����~�)xc-�az=�m'E��l��'�!���:�L���!u	y��âCe,�7|��s����Ӯ�ǌ}��4O&��E���K�1b�Fnl�F5	Y���e��ѵ���#E�6YF|~J����]&��f�A��\�h,j���ҥ��-��wcYS��I0���+�L�h6��1��p�S� FHx�^Jb75"��֣YAj����2�N�E\��$,Dd�<V���w����<g��ߣ��[����خV�� B��         �   x�e�;�0  й=s���(�jm#�K��6PT�pzuvy����iU�s��*Q������Mc?Ҕk�~0E�W������An3�^{j���]+��N=rD�i#k�}��l��*�f�ha��ɝdr�$~@c2e�\EX��"�g6ޫ�9HxE�>�5�     
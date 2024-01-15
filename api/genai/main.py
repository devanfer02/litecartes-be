import google.generativeai as genai 
import os
from IPython.display import Markdown

gemini_api_key = os.environ.get('GEMINI_API_KEY')
genai.configure(api_key = gemini_api_key)

model = genai.GenerativeModel('gemini-pro')
response = model.generate_content(
    '''
    From this literacy text : 
    
    Artificial Intelligence (AI) menjadi semakin terdepan dalam dunia teknologi. 
    Kemampuannya untuk mengenali pola, belajar dari pengalaman, dan membuat keputusan semakin kompleks, 
    memperluas penggunaannya dari industri hingga ke kehidupan sehari-hari. 
    Dari asisten virtual yang bisa merespon pertanyaan kita hingga mobil otonom yang dapat mengemudi sendiri, 
    kehadiran AI telah mengubah cara kita berinteraksi dengan teknologi.
    Menurutmu, apakah dampak ini bisa membawa positif terhadap perkembangan manusia atau malah sebaliknya ?

    With this response to it : 
    Ya, dengan bantuan AI bisa membantu orang mendapatkan akses lebih mudah ke sumber belajar

    Give me a rating from 0 to 10 with how much the response related to literacy text and question in percent
    '''
)

Markdown(response.txt)


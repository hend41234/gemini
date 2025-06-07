curl https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-lite:generateContent?key=$GEMINI_API_KEY \
  -H 'Content-Type: application/json' \
  -X POST \
  -d '{
    "contents": [
      {
        "parts": [
          {
            "text": "buatkan beberapa paragraf, yang menceritakan kontradiktif fakta, atau plotwist dunia yang jarang diketahui manusia, tulis dalam bahasa inggris. hanya tulis pragraf nya saja, tidak perlu ditambahin apa apa. hanya jawabannya saja."
          }
        ]
      }
    ],
    "generationConfig": {
      "stopSequences": [
        "Title"
      ],
      "temperature": 1.0,
      "maxOutputTokens": 800,
      "topP": 0.8,
      "topK": 10
    }
  }'
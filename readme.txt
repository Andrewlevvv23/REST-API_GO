"math/rand" - генерація рандомного числа 
              number := rand.Intn(100)

"time" - робота з таймаутами, часом
         time.Sleep(500 * time.Millisecond)  
         time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second()           

"defer" - відкладений виклик функції в стеку. Аналог finally(), First In - Last Out

"& *" - вказівники, для позбавлення від копіювання великих обʼєктів/змінних.
        Роблять змінні з посиланням на обʼєкт, зберігаються в heap


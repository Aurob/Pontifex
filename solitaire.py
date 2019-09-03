import random
def (deck):
    if deck.index('a') == len(deck)-1:
        deck.insert(0,deck.pop())
        
    elif deck.index('a') == len(deck)-2:
        deck.insert(len(deck)-1,deck.pop(deck.index('a')))
        
    else:
        deck.insert(deck.index('a')+1,deck.pop(deck.index('a')))
        
    if deck.index('b') == len(deck)-1:
        deck.insert(2,deck.pop())
        
    elif deck.index('b') == len(deck)-2:
        deck.insert(1,deck.pop(deck.index('b')))
        
    else:
        deck.insert(deck.index('b')+2,deck.pop(deck.index('b')))

    if deck.index('a') < deck.index('b'):
        right = deck[deck.index('b')+1:]
        left = deck[:deck.index('a')]
        middle = deck[deck.index('a')+1:deck.index('b')]
        deck = right +["a"]+ middle +["b"]+ left

    else:
        right = deck[deck.index('a')+1:]
        left = deck[:deck.index('b')]
        middle = deck[deck.index('b')+1:deck.index('a')]
        deck = right +["b"]+ middle +["a"]+ left

    if deck[-1] in ['a','b']:
        right = deck[53:]
        middle = deck[:53]
        left = deck[53:-1]
        
    else:
        right = deck[-1:]
        middle = deck[:deck[-1]]
        left = deck[deck[-1]:-1]

    deck = left+middle+right

    if deck[0] == 'a' or deck[0] == 'b':
        return [deck,deck[53]]
    else:
        return [deck,deck[deck[0]]]

def encrypt(deck,text,o):
    to = []
    for a in text:
        #print(a)
        dkey2 = step1(deck)
        while True:
            if dkey2[1] in ['a','b']:
                dkey2 = step1(dkey2[0])
            else:
                break

        if o ==2:
            if abc.index(a) < dkey2[1]:
                to.append(abc[abc.index(a)+26 - dkey2[1]])
            else:
                to.append(abc[abc.index(a) - dkey2[1]])

                
        if o == 1:
            to.append(abc[(((dkey2[1]%26)-1)+abc.index(a.lower())+1)%26])
        deck = dkey2[0]
    return to

text = input("Insert text to encrypt: ").replace(' ','')
abc = 'abcdefghijklmnopqrstuvwxyz'   

dkey = list(i for i in range(1,53))
dkey.extend(('a','b'))

random.shuffle(dkey)

print("The deck order is: ", dkey)

encrypted = encrypt(dkey.copy(),text,1)

decrypted = encrypt(dkey.copy(),encrypted,2)

print("Encrypted text: ",encrypted)


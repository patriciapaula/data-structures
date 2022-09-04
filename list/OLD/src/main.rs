/*struct Memory<T> {
    buffer: Vec<T>,
    free: Vec<usize>,
}

impl<T> Memory<T> {
    fn new() -> Self {
        Self {
            buffer: Vec::<T>::new(),
            free: Vec::new(),
        }
    }
    
    fn alloc(&mut self, value: T) -> usize {
        if let Some(index) = self.free.pop() {
            self.buffer[index] = value;
            index
        } else {
            self.buffer.push(value);
            self.buffer.len() - 1
        }
    }

    fn dealloc(&mut self, ptr: usize) {
        self.free.push(ptr);
    }

    fn deref(&self, ptr: usize) -> &T {
        &self.buffer[ptr]
    }

    fn deref_mut(&mut self, ptr: usize) -> &mut T {
        &mut self.buffer[ptr]
    }
} */


use std::ptr;
use std::marker::PhantomData;

#[derive(Clone)]
struct Lista {
    info: Option<usize>,
    prox: PhantomData,
}

impl Lista {
    fn cria() -> Self {
        Self {
            info: None,
            prox: *mut Option<usize>,
        }
    }

    fn insere(lista: Lista, info: Option<usize>) {
        let novo = Self::cria();
        novo.info = info;
        novo.prox = *lista;
    }

    //fn insere(&mut self, memory: &mut Memory<Lista>, info: Option<usize>) {
    //    let novo = memory.alloc(Lista::cria());
    //    self.
    //}
}


fn main() {
    println!("Hello, world!");
}

package com.dresser.infrastructure.db;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Collectors;

public class MemoryDB implements DB {
    
    private final Map<String, Map<Integer, Object>> storage = new ConcurrentHashMap<>();
    private final Map<String, Integer> counters = new ConcurrentHashMap<>();
    
    @Override
    public Object create(String collection, int id, Object entity) {
        storage.computeIfAbsent(collection, k -> new ConcurrentHashMap<>())
               .put(id, entity);
        return entity;
    }
    
    @Override
    public Object findById(String collection, Object id) {
        Map<Integer, Object> collectionMap = storage.get(collection);
        if (collectionMap == null) {
            return null;
        }
        return collectionMap.get(id);
    }
    
    @Override
    public List<Object> findAll(String collection) {
        Map<Integer, Object> collectionMap = storage.get(collection);
        if (collectionMap == null) {
            return new ArrayList<>();
        }
        return new ArrayList<>(collectionMap.values());
    }
    
    @Override
    public List<Object> findBy(String collection, Map<String, Object> filter) {
        // Simple implementation - this would need more sophisticated handling in a real app
        return findAll(collection);
    }
    
    @Override
    public void update(String collection, Object id, Object entity) {
        Map<Integer, Object> collectionMap = storage.get(collection);
        if (collectionMap != null) {
            collectionMap.put((Integer) id, entity);
        }
    }
    
    @Override
    public void delete(String collection, Object id) {
        Map<Integer, Object> collectionMap = storage.get(collection);
        if (collectionMap != null) {
            collectionMap.remove(id);
        }
    }
    
    @Override
    public int nextId(String collection) {
        return counters.compute(collection, (k, v) -> v == null ? 1 : v + 1);
    }
    
    @Override
    public void close() {
        // Nothing to close for in-memory DB
    }
} 
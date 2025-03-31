package com.dresser.infrastructure.db;

import java.util.List;
import java.util.Map;

public interface DB {
    
    /**
     * Creates a new entity in the database
     * 
     * @param collection The collection name
     * @param id The entity ID
     * @param entity The entity to create
     * @return The created entity
     */
    Object create(String collection, int id, Object entity);
    
    /**
     * Finds an entity by ID
     * 
     * @param collection The collection name
     * @param id The entity ID
     * @return The found entity or null
     */
    Object findById(String collection, Object id);
    
    /**
     * Finds all entities in a collection
     * 
     * @param collection The collection name
     * @return List of entities
     */
    List<Object> findAll(String collection);
    
    /**
     * Finds entities by a filter
     * 
     * @param collection The collection name
     * @param filter The filter criteria
     * @return List of matching entities
     */
    List<Object> findBy(String collection, Map<String, Object> filter);
    
    /**
     * Updates an entity
     * 
     * @param collection The collection name
     * @param id The entity ID
     * @param entity The updated entity
     */
    void update(String collection, Object id, Object entity);
    
    /**
     * Deletes an entity
     * 
     * @param collection The collection name
     * @param id The entity ID
     */
    void delete(String collection, Object id);
    
    /**
     * Generates the next available ID for a collection
     * 
     * @param collection The collection name
     * @return The next ID
     */
    int nextId(String collection);
    
    /**
     * Closes the database connection
     */
    void close();
} 
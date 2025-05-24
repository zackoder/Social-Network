"use client";

import { useState, useEffect, useRef } from 'react';
import styles from './notificationDropdown.module.css';

export default function NotificationDropdown({ isOpen, onClose }) {
  const [notifications, setNotifications] = useState([]);
  const [loading, setLoading] = useState(false);
  const [offset, setOffset] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const dropdownRef = useRef(null);

  // Fetch notifications
  const fetchNotifications = async () => {
    if (loading || !hasMore) return;
    
    try {
      setLoading(true);
      const response = await fetch(`/api/notifications?limit=5&offset=${offset}`);
      
      if (!response.ok) {
        // throw new Error('Failed to fetch notifications');
      }
      
      const data = await response.json();
      
      if (offset === 0) {
        setNotifications(data.notifications);
      } else {
        setNotifications(prev => [...prev, ...data.notifications]);
      }
      
      setHasMore(data.hasMore);
      setOffset(prev => prev + data.notifications.length);
    } catch (error) {
      console.error('Error fetching notifications:', error);
    } finally {
      setLoading(false);
    }
  };

  // Load more notifications
  const handleLoadMore = () => {
    fetchNotifications();
  };

  // Close dropdown when clicking outside
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        onClose();
      }
    };

    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    }
    
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [isOpen, onClose]);

  // Fetch notifications when dropdown is opened
  useEffect(() => {
    if (isOpen) {
      setOffset(0);
      setHasMore(true);
      fetchNotifications();
    }
  }, [isOpen]);

  if (!isOpen) return null;

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  return (
    <div className={styles.dropdown} ref={dropdownRef}>
      <div className={styles.header}>
        <h3>Notifications</h3>
      </div>
      
      <div className={styles.notificationsList}>
        {notifications.length === 0 && !loading ? (
          <div className={styles.emptyState}>No notifications</div>
        ) : (
          <>
            {notifications.map((notification) => (
              <div key={notification.id} className={styles.notificationItem}>
                <div className={styles.message}>{notification.message}</div>
                <div className={styles.time}>{formatDate(notification.created_at)}</div>
              </div>
            ))}
            
            {hasMore && (
              <button 
                className={styles.loadMoreBtn} 
                onClick={handleLoadMore}
                disabled={loading}
              >
                {loading ? 'Loading...' : 'Load More'}
              </button>
            )}
          </>
        )}
      </div>
    </div>
  );
}

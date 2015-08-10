//
//  KBLaunchService.h
//  Keybase
//
//  Created by Gabriel on 3/12/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import <Foundation/Foundation.h>

#import "KBInstallable.h"
#import "KBComponent.h"

@interface KBLaunchService : NSObject

@property (readonly) NSString *label;
@property (readonly) NSDictionary *plist;
@property (readonly) NSString *versionPath;
@property (readonly) NSString *logFile;
@property (readonly) NSString *bundleVersion;

@property (readonly) KBComponentStatus *componentStatus;

- (instancetype)initWithLabel:(NSString *)label bundleVersion:(NSString *)bundleVersion versionPath:(NSString *)versionPath plist:(NSDictionary *)plist logFile:(NSString *)logFile;

- (NSString *)plistDestination;

- (void)install:(KBCompletion)completion;
- (void)uninstall:(KBCompletion)completion;
- (void)start:(KBCompletion)completion;
- (void)stop:(KBCompletion)completion;

- (GHODictionary *)componentStatusInfo;

- (void)updateComponentStatus:(NSTimeInterval)timeout completion:(KBCompletion)completion;
- (void)refreshLaunchStatus:(KBCompletion)completion;

@end
